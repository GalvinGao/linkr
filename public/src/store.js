import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const emptyCredentials = {
  loggedIn: false,
  username: "",
  token: ""
}

export default new Vuex.Store({
  state: {
    credentials: emptyCredentials
  },
  mutations: {
    updateAuth(state, newState) {
      state.credentials = newState
    },
    logout(state) {
      state.credentials = emptyCredentials
    }
  },
  actions: {

  }
})
