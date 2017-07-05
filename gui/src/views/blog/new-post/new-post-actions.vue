<template>
  <div id="new-post-actions">
    <template v-if="authInfo.role === roles.Admin">
      <ui-icon-button
        type="secondary"
        color="white"
        icon="arrow_back"
        @click="$router.push({name: 'Blog.Posts'})">
      </ui-icon-button>

      <ui-icon-button
        type="secondary"
        color="white"
        icon="save"
        @click="save">
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
    }),

    // communicate with content
    in ({cmd, data}) {
      if (cmd === 'savePost') {
        this.savePost(data)
      }
    },

    save () {
      this.$emit('channel', {cmd: 'savePost'})
    },
    savePost (data) {
      console.log(data.content)
    }
  },

  mounted () {
    this.getAuthInfo({}) // will be deprecated
  }
}
</script>

<style lang="sass">
#new-post-actions
  padding-right: 16px
  height: 36px
  .ui-icon-button
    width: 36px
    height: 36px
</style>
