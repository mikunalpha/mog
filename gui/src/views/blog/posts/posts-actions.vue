<template>
  <div id="posts-actions">
    <ui-textbox
      placeholder="Search ..."
      v-model="keyword">
    </ui-textbox>
    <ui-icon-button
      type="secondary"
      color="white"
      icon="search"
      @click="search">
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
  data () {
    return {
      keyword: ''
    }
  },

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

    search () {
      this.$emit('channel', {cmd: 'searchPosts', data: this.keyword.trim()})
    }
  },

  mounted () {
    //
  }
}
</script>

<style lang="sass" scoped>
#posts-actions
  padding-right: 16px
  height: 36px
  .ui-textbox
    display: inline-block
  .ui-icon-button
    width: 36px
    height: 36px
</style>

<style lang="sass">
#posts-actions
  .ui-textbox
    input
      padding: 0 5px
      letter-spacing: 1px
      color: #ffffff
      border-bottom-color: rgba(255,255,255,.5)
</style>
