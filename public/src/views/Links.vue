<template>
  <div class="container links">
    <h1>
      <a-icon type="link"/>
      Links
    </h1>
    <a-table :columns="columns"
             :rowKey="record => record.link_id"
             :dataSource="data"
             :pagination="pagination"
             :loading="loading"
             bordered
             @change="handleTableChange"
             class="link-table"
    >
      <template slot="link_id" slot-scope="link_id">
        <span class="monospaced">
          {{ link_id }}
        </span>
      </template>
      <template slot="created_at" slot-scope="created_at">
        <a-tooltip placement="bottomLeft" style="cursor: default">
          <template slot='title'>
            {{ created_at.raw }}
          </template>
          <a-icon type="clock-circle"/>
          {{ created_at.fromNow }}
        </a-tooltip>
      </template>
      <template slot="updated_at" slot-scope="updated_at">
        <a-tooltip placement="bottomLeft" style="cursor: default">
          <template slot='title'>
            {{ updated_at.raw }}
          </template>
          <a-icon type="clock-circle"/>
          {{ updated_at.fromNow }}
        </a-tooltip>
      </template>
      <template slot="short_url" slot-scope="record, short_url">
        <span class="monospaced" v-line-clamp="2">
          <a-tooltip style="display: inline">
            <template slot='title'>
              Copy short link to Clipboard
            </template>
            <a @click="() => copyLink(record.short_url.full)" style="cursor: pointer">{{short_url.path}}</a>
          </a-tooltip>
        </span>
      </template>
      <template slot="long_url" slot-scope="long_url" class="monospaced">
        <span target="_blank" class="monospaced long-url" v-line-clamp="3">{{long_url}}</span>
      </template>
      <template slot="action" slot-scope="record">
        <a-button size="small" icon="edit">Edit</a-button>
        <a-divider type="vertical"/>
        <a-popconfirm title="Delete Link?" @confirm="() => confirmDelete(record.link_id)" @cancel="cancel" okText="Delete" okType="danger" placement="topRight">
          <a-icon slot="icon" type="question-circle-o" style="color: #ff4d4f"/>
          <a-button type="danger" size="small" icon="delete">Delete</a-button>
        </a-popconfirm>
      </template>
    </a-table>
  </div>
</template>

<script>
  const columns = [{
    title: '#',
    dataIndex: 'link_id',
    sorter: true,
    width: '3%',
    scopedSlots: {customRender: 'link_id'},
  }, {
    title: 'Created',
    dataIndex: 'created_at',
    sorter: true,
    width: '15%',
    scopedSlots: {customRender: 'created_at'},
  }, {
    title: 'Last Modified',
    dataIndex: 'updated_at',
    sorter: true,
    width: '15%',
    scopedSlots: {customRender: 'updated_at'},
  }, {
    title: 'Shortened',
    dataIndex: 'short_url',
    width: '5%',
    scopedSlots: {customRender: 'short_url'},
  }, {
    title: 'Original',
    dataIndex: 'long_url',
    width: '37%',
    scopedSlots: {customRender: 'long_url'},
  }, {
    title: 'Action',
    width: '20%',
    scopedSlots: {customRender: 'action'},
  }];

  export default {
    name: "Links",
    data() {
      return {
        data: [],
        pagination: {},
        loading: false,
        columns
      }
    },
    mounted() {
      this.fetch()
    },
    methods: {
      handleTableChange(pagination, filters, sorter) {
        const pager = {...this.pagination};
        pager.current = pagination.current;
        this.pagination = pager;

        let params = {
          limits: pagination.pageSize,
          page: pagination.current,
          sort_field: sorter.field,
          sort_order: sorter.order
        }
        console.log("params:", params)
        this.fetch(params)
      },
      fetch(params = {}) {
        this.loading = true
        console.log("params:", {
          limits: 10,
          page: 1,
          sort_field: params.sort_field,
          sort_order: params.sort_order,
          ...params
        })
        this.$http.get('/api/link', {
          params: {
            limits: 10,
            page: 1,
            ...params
          },
          headers: {
            'Authorization': 'Bearer ' + this.$store.state.credentials.token
          }
        })
          .then((resp) => {
            let data = resp.data
            const pagination = {...this.pagination}
            // Read total count from server
            // pagination.total = data.totalCount;
            pagination.total = data.total_record
            data.records.forEach((el) => {
              el.created_at = {
                "raw": this.$moment(el.created_at).format(),
                "fromNow": this.$moment(el.created_at).fromNow()
              }
              el.updated_at = {
                "raw": this.$moment(el.updated_at).format(),
                "fromNow": this.$moment(el.updated_at).fromNow()
              }
              el.short_url = {
                "full": [window.location.origin, el.short_url].join("/"),
                "path": el.short_url
              }
            })
            this.data = data.records
            this.pagination = pagination
          })
          .finally(() => {
            this.loading = false
          })
      },
      copyLink(linkContent) {

      },
      confirmDelete(linkId) {

      }
    }
  }
</script>

<style scoped>
  .monospaced {
    font-family: Consolas, Courier, Courier New, monospace;
  }

  .long-url {
    line-break: anywhere;
    word-break: break-all;
  }

  .link-table {
    min-width: 720px;
  }
</style>
