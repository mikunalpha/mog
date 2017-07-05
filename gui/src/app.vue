<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app'
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
