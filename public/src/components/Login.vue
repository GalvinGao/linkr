<template>
  <a-modal
    title="Login"
    centered
    :closable="false"
    :keyboard="false"
    :maskClosable="false"
    :destroyOnClose="true"
    :maskStyle="{'backdrop-filter': 'blur(5px)', 'background-image': 'radial-gradient(ellipse at center, #333 0%, black 100%)'}"
    v-model="needLogin"
  >
    <template slot="footer">
      <a-button key="submit" type="primary" :loading="ajaxStatus" @click="startAuthentication">
        Submit
      </a-button>
    </template>
    <a-form
    id="loginForm"
    :form="form"
    class="login-form"
    @submit="startAuthentication"
  >
    <a-form-item>
      <a-input
        v-decorator="[
          'username',
          { rules: [{ required: true, message: 'Please input your username!' }] }
        ]"
        placeholder="Username"
      >
        <a-icon
          slot="prefix"
          type="user"
          style="color: rgba(0,0,0,.25)"
        />
      </a-input>
    </a-form-item>
    <a-form-item>
      <a-input
        v-decorator="[
          'password',
          { rules: [{ required: true, message: 'Please input your Password!' }] }
        ]"
        type="password"
        placeholder="Password"
      >
        <a-icon
          slot="prefix"
          type="lock"
          style="color: rgba(0,0,0,.25)"
        />
      </a-input>
    </a-form-item>
  </a-form>
  </a-modal>
</template>

<script>
  export default {
    name: "Login",
    data() {
      return {
        needLogin: !this.$store.state.authentication.loggedIn,
        ajaxStatus: false,
        form: {}
      }
    },
    methods: {
      startAuthentication() {
        let fields = this.form.getFieldsValue(['username', 'password'])
        let preparedCredentials = this.$prepareCredentials(fields.password)
        let data = {
          username: fields.username,
          key: preparedCredentials.key,
          password: preparedCredentials.encryptedPassword
        }
        let postData = this.$queryString.stringify(data)

        this.$http.post("/admin/api/login", postData)
          .then((resp) => {
            console.log(resp.data.toJSON())
          })
          .catch(function (error) {
            console.error(error);
          });
      }
    },
    beforeMount () {
      this.form = this.$form.createForm(this)
    }
  }
</script>

<style scoped>
#loginForm .ant-form-item {
  margin-bottom: 4px;
}
</style>
