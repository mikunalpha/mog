<template>
  <div id="posts-actions">
    <ui-icon-button
      type="secondary"
      color="white"
      icon="search">
    </ui-icon-button>
    <template v-if="authInfo.role === roles.Admin">
      <ui-icon-button
        type="secondary"
        color="white"
        icon="add"
        @click="$router.push({name: 'Blog.NewPost'})">
      </ui-icon-button>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  computed: {
    ...mapGetters({
      authInfo: 'authInfo',
      roles: 'roles'
    })
  },

  methods: {
    ...mapActions({
      getAuthInfo: 'getAuthInfo'
    })
  },

  mounted () {
    this.getAuthInfo({}) // will be deprecated
  }
}
</script>

<style lang="sass" scoped>
#posts-actions
  padding-right: 16px
  height: 36px
  .ui-icon-button
    width: 36px
    height: 36px
</style>
