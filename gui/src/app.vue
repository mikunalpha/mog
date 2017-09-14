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
      getAuthInfo: 'getAuthInfo',
      getStatus: 'getStatus'
    })
  },

  created () {
    // Get authentication information from server and put into vuex authInfo.
    this.getAuthInfo({})

    // Get status from server and put into vuex status.
    this.getStatus({
      successCallback: (status) => {
        if (status.administered === false) {
          this.$router.replace({name: 'Auth.NewAdmin'})
          return
        }
      },
      errorCallback: ({status, error}) => {
        console.log(status)
        this.$router.replace({name: '500'})
      }
    })
  }
}
</script>

<style lang="sass">
@import '~@/assets/css/normalize.css'
@import '~@/assets/css/style.css'
@import '~@/assets/sass/variables.sass'
@import '~@/assets/css/keen-ui.min.css'

html, body
  min-height: 100vh
  margin: 0
  padding: 0
  width: 100%

#app
  min-height: 100vh
  font-family: 'Robot'
  color: $fontColor
</style>
