<template>
  <div id="post-actions">
    <template v-if="mode === 'view' && authInfo.role === roles.Admin">
      <ui-icon-button
        type="secondary"
        color="white"
        icon="arrow_back"
        @click="$router.push({name: 'Blog.Posts'})">
      </ui-icon-button>

      <ui-icon-button
        type="secondary"
        color="white"
        icon="edit"
        @click="edit">
      </ui-icon-button>

      <ui-icon-button
        type="secondary"
        color="white"
        icon="delete"
        @click="deleteConfirm">
      </ui-icon-button>

      <ui-confirm
        class="deleteConfirm"
        ref="deleteConfirm"
        title="Delete Post"
        type="danger"
        confirmButtonText="Delete"
        denyButtonText="Keep"
        dismissOn="backdrop esc"
        @confirm=""
        @deny="">
        Are you sure you want to delete the post?
      </ui-confirm>
    </template>

    <template v-else-if="mode === 'edit' && authInfo.role === roles.Admin">
      <ui-icon-button
        type="secondary"
        color="white"
        icon="arrow_back"
        @click="view">
      </ui-icon-button>

      <ui-icon-button
        type="secondary"
        color="white"
        icon="save"
        @click="save">
      </ui-icon-button>
    </template>

    <template v-else>
      <ui-icon-button
        type="secondary"
        color="white"
        icon="arrow_back"
        @click="$router.push({name: 'Blog.Posts'})">
      </ui-icon-button>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  data () {
    return {
      mode: 'view'
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

    // communicate with content
    in ({cmd, data}) {
      if (cmd === 'savePost') {
        this.savePost(data)
      }
    },

    // Change to edit mode
    deleteConfirm () {
      this.$refs.deleteConfirm.open()
    },
    edit () {
      this.$emit('channel', {cmd: 'changeMode', data: 'edit'})
      this.mode = 'edit'
    },
    view () {
      this.$emit('channel', {cmd: 'changeMode', data: 'view'})
      this.mode = 'view'
    },
    save () {
      this.$emit('channel', {cmd: 'savePost'})
    },
    savePost (data) {
      console.log(data.content)
    }
  },

  // beforeRouteEnter (to, from, next) {
    // if (from.name === 'Blog.Posts') {
    //   console.log(from)
    //   to.meta.backToRoute = {name: 'Blog.Posts'}
    // }
    // next()
  // },

  mounted () {
    this.getAuthInfo({}) // will be deprecated
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#post-actions
  padding-right: 16px
  height: 36px
  .ui-icon-button
    width: 36px
    height: 36px
  .deleteConfirm
    color: $fontColor
</style>
