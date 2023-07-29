import Vue from 'vue'
import Vuex from 'vuex'
import Cookies from 'js-cookie'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    endpoints: {
      endpoint1: Cookies.get('endpoint1') || '',
      endpoint2: Cookies.get('endpoint2') || '',
    },
    interval: Cookies.get('interval') || 500
  },
  mutations: {
    setEndpoints(state, endpoints) {
      state.endpoints = endpoints
    },
    setInterval(state, interval) {
      state.interval = interval
    }
  },
  actions: {
    saveEndpoints({commit}, endpoints) {
      Cookies.set('endpoint1', endpoints.endpoint1);
      if (endpoints.endpoint2) {
        Cookies.set('endpoint2', endpoints.endpoint2);
      } else {
        Cookies.remove('endpoint2');
      }
      commit('setEndpoints', endpoints)
    },
    saveInterval({commit}, interval) {
      Cookies.set('interval', interval)
      commit('setInterval', interval)
    }
  }
})
