<template>
  <div id="app">
    <Login/>
    <a-layout id="components-layout-demo-top" class="layout" v-if="this.$store.state.credentials.loggedIn">
      <a-layout-header>
        <div class="logo">
          LINKR
        </div>
        <a-menu
          theme="dark"
          mode="horizontal"
          :defaultSelectedKeys="['nav:dashboard']"
          :style="{ lineHeight: '64px' }"
        >
          <a-menu-item key="nav:links">
            <router-link to="/links"><a-icon type="link" /> Links</router-link>
          </a-menu-item>
          <a-menu-item key="nav:settings" to="/settings">
            <router-link to="/settings"><a-icon type="setting" /> Settings</router-link>
          </a-menu-item>
          <a-button type="danger" shape="circle" icon="logout" @click="() => this.$store.commit('logout')"
                    class="logout-button">
            <SROnly text="Logout"/>
          </a-button>
        </a-menu>
      </a-layout-header>
      <a-layout-content style="padding: 0 50px">
        <a-breadcrumb style="margin: 16px 0" :routes="routes">
          <template slot="itemRender" slot-scope="{route, params, routes, paths}">
            <span v-if="routes.indexOf(route) === routes.length - 1">
              {{route.name}}
            </span>
            <router-link v-else :to="`${basePath}/${paths.join('/')}`">
              {{route.name}}
            </router-link>
          </template>
        </a-breadcrumb>
        <div class="router-root">
          <router-view/>
        </div>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Linkr {{ this.currentYear }}. Source code available at <a href="https://github.com/GalvinGao/linkr"
                                                                  target="_blank">GitHub</a>.
      </a-layout-footer>
    </a-layout>
  </div>
</template>

<script>
  import Login from '@/components/Login'
  import SROnly from '@/components/sr-only'

  export default {
    components: {
      SROnly,
      Login
    },
    mounted() {
      this.updateBreadcrumb()
    },
    data() {
      return {
        routes: []
      }
    },
    methods: {
      updateBreadcrumb() {
        let breadcrumbs = [];
        for (let segment of this.$router.currentRoute.matched) {
          breadcrumbs.push({
            "name": segment.name,
            "path": segment.path
          })
        }
        this.routes = breadcrumbs
      }
    },
    watch: {
      '$route': ['updateBreadcrumb']
    },
    computed: {
      currentYear() {
        return new Date().getFullYear()
      }
    }
  }
</script>

<style>
  #components-layout-demo-top .logo {
    width: 120px;
    height: 31px;
    line-height: 31px;
    margin: 16px 24px 16px 0;
    float: left;
    text-align: center;
    color: #cccccc;
    letter-spacing: 1px;
    font-weight: bold;
    cursor: default;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .logout-button {
    float: right;
    margin: 16px;
  }

  .logout-button > .anticon {
    margin-right: 0 !important;
  }

  .router-root {
    background: #fff;
    min-height: calc(100vh - 120px);
    width: 100%;

    overflow-x: scroll;
  }
  .container {
    padding: 2rem 1rem;
    margin-right: auto;
    margin-left: auto;
  }

  @media (min-width: 576px) {
    .container {
      max-width: 540px;
    }
  }

  @media (min-width: 768px) {
    .container {
      max-width: 720px;
    }
  }

  @media (min-width: 992px) {
    .container {
      max-width: 960px;
    }
  }

  @media (min-width: 1200px) {
    .container {
      max-width: 1140px;
    }
  }

  @media (min-width: 1536px) {
    .container {
      max-width: 1450px;
    }
  }
</style>
