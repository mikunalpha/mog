<template>
  <div id="post-actions">
    <template v-if="mode === 'view' && authInfo.role === roles.Admin">
      <ui-icon-button
        type="secondary"
        color="white"
        icon="arrow_back"
        @click="backToPostsView">
      </ui-icon-button>

      <ui-icon-button
        v-if="post.data.id"
        type="secondary"
        color="white"
        icon="edit"
        @click="edit">
      </ui-icon-button>

      <ui-icon-button
        v-if="post.data.id"
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
        @confirm="destroy"
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
        :loading="updating"
        @click="update">
      </ui-icon-button>
    </template>

    <template v-else>
      <ui-icon-button
        type="secondary"
        color="white"
        icon="arrow_back"
        @click="backToPostsView">
      </ui-icon-button>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  data () {
    return {
      mode: 'view',
      updating: false
    }
  },

  computed: {
    ...mapGetters({
      authInfo: 'authInfo',
      roles: 'roles',
      post: 'post'
    })
  },

  methods: {
    ...mapActions({
      getAuthInfo: 'getAuthInfo',
      getStatus: 'getStatus',
      updatePost: 'updatePost',
      deletePost: 'deletePost'
    }),

    // communicate with content
    in ({cmd, data}) {
      if (cmd === 'updatePost') {
        this.updateEditedPost(data)
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

    // Update post
    update () {
      this.updating = true
      this.$emit('channel', {cmd: 'updatePost'})
    },
    updateEditedPost (editedPost) {
      let t = this

      t.updatePost({
        id: t.$route.params.id,
        post: editedPost,
        success: (post) => {
          t.$emit('channel', {cmd: 'refreshPost'})
          this.updating = false
        },
        error: (status, e) => {
          this.updating = false
        }
      })
    },

    // Delete post
    destroy () {
      this.deletePost({
        id: this.$route.params.id,
        success: (post) => {
          this.getStatus({})
          this.$router.push({name: 'Blog.Posts'})
        },
        error: (status, e) => {
        }
      })
    },

    backToPostsView () {
      this.$router.push({name: 'Blog.Posts', query: {page: 1}})
    }
  },

  // beforeRouteEnter (to, from, next) {
    // if (from.name === 'Blog.Posts') {
    //   console.log(from)
    //   to.meta.backToRoute = {name: 'Blog.Posts'}
    // }
    // next()
  // },

  created () {
    // this.getAuthInfo({}) // will be deprecated
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
