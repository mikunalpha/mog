<template>
  <div id="auth-view">
    <template v-if="status.gotAt !== null">
    <router-view></router-view>
    </template>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'auth-view',

  computed: {
    ...mapGetters({
      status: 'status'
    })
  }
}
</script>

<style lang="sass" scoped>
#auth-view
</style>
