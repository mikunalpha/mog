<template>
  <div id="auth-view">
    <template v-if="hasGotStatus">
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
    }),

    hasGotStatus () {
      return this.status.gotAt !== null
    }
  }
}
</script>
