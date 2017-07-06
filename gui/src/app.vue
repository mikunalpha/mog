<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'app',

  methods: {
    ...mapActions({
      getStatus: 'getStatus',
      getAuthInfo: 'getAuthInfo'
    })
  },

  created () {
    this.getAuthInfo({
      // success: (info) => {
      //   console.log(info)
      //   next()
      // },
      // error: (status, e) => {
      //   console.log(status)
      //   next()
      // }
    })

    this.getStatus({
      success: (status) => {
        if (status.administered === false) {
          this.$router.replace({name: 'Auth.NewAdmin'})
          return
        }
      },
      error: (status, e) => {
        console.log(status)
        this.$router.replace({name: '500'})
      }
    })
  }
}
</script>

<style lang="sass">
@import './assets/css/normalize.css'
@import './assets/css/style.css'
@import './assets/css/keen-ui.min.css'

#app
  min-height: 100vh
  font-family: 'Robot'
</style>
