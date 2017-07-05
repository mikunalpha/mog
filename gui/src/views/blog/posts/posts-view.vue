<template>
  <div id="posts-view">
    <div class="post" v-for="p in posts">
      <div class="title">
        <span @click="goToViewPost({id: p.id})">{{ p.title }}</span>
      </div>

      <div class="created-at">
        <ui-icon class="time" icon="date_range"></ui-icon>
        {{ p.created_at | toDatetime }}
      </div>

      <div class="content" v-html="p.content">
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
// import NewPostView from '@/views/admin/posts/new-post/new-post-view'
// import EditPostView from '@/views/admin/posts/edit-post/edit-post-view'

export default {
  data () {
    return {
      // newPostViewIsOpen: false
      // editPostViewIsOpen: false
    }
  },

  computed: {
    ...mapGetters({
      roles: 'roles',
      posts: 'posts',
      post: 'post'
    })
  },

  components: {
    // NewPostView,
    // EditPostView
  },

  methods: {
    ...mapActions({
      getPosts: 'getPosts',
      getPost: 'getPost'
    }),
    goToViewPost ({id}) {
      this.$router.push({name: 'Blog.Post', params: {id}})
    }
  },

  mounted () {
    this.getPosts({})
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#posts-view
  position: relative
  padding-top: 20px
  padding-bottom: 120px
  min-height: 100vh
  background-color: #f0f0f0
  .post
    position: relative
    margin: 0 auto
    margin-bottom: 20px
    padding: 30px 20px
    max-width: 940px
    background-color: #ffffff
    box-shadow: 0 1px 0 #e0e0e0
    .title
      color: $fontColor
      font-size: 1.6rem
      span
        cursor: pointer
    .created-at
      padding: 10px 0
      font-size: .9rem
      color: lighten($fontColor, 10%)
      .time
        margin-top: -3px
        font-size: 1rem
    .content
        letter-spacing: 1px
        line-height: 200%
</style>
