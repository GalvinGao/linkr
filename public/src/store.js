import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    authentication: {
      loggedIn: false,
      username: "",
      token: ""
    }
  },
  mutations: {
    updateAuthenticationState(state) {
      this.state.authentication = state
    }
  },
  actions: {

  }
})
