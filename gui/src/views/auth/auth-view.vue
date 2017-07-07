<template>
  <div id="auth-view">
<<<<<<< HEAD
    <template v-if="status.gotAt !== null">
    <router-view></router-view>
    </template>
=======
    <router-view></router-view>
>>>>>>> da1474513963a90bd1176e353ded31eeb3c42dfc
  </div>
</template>

<script>
<<<<<<< HEAD
import { mapGetters } from 'vuex'

export default {
  name: 'auth-view',

  computed: {
    ...mapGetters({
      status: 'status'
    })
  }
=======
export default {
>>>>>>> da1474513963a90bd1176e353ded31eeb3c42dfc
}
</script>

<style lang="sass" scoped>
#auth-view
</style>
