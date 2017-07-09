<template>
  <div id="posts-view">
    <template v-if="posts.gotAt !== null && posts.data.length > 0">
      <div class="post" :class="{unpublished: !p.published}" v-for="p in posts.data">
        <div class="title">
          <!-- <span @click="goToViewPost({id: p.id})">{{ p.title }}</span> -->
          <router-link 
            tag="a"
            :to="{name: 'Blog.Post', params: {id: p.id}}">
            {{ p.title }}
          </router-link>
        </div>

        <div class="created-at">
          <ui-icon class="time" icon="date_range"></ui-icon>
          {{ p.created_at | toDatetime }}
        </div>

        <div class="content" v-html="p.content">
        </div>
      </div>
    </template>
    <template v-else-if="posts.gotAt !== null && posts.data.length === 0">
      <div class="no-posts">
        No Posts
      </div>
    </template>
    <template v-else>
      <div class="loading-posts">
        Loading...
      </div>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  data () {
    return {
    }
  },

  computed: {
    ...mapGetters({
      roles: 'roles',
      posts: 'posts'
    })
  },

  components: {
  },

  methods: {
    ...mapActions({
      getPosts: 'getPosts'
    }),
    goToViewPost ({id}) {
      this.$router.push({name: 'Blog.Post', params: {id}})
    }
  },

  mounted () {
    this.getPosts({options: {
      limit: 2,
      excerpt: 1
    }})
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
    padding: 30px 20px 40px 20px
    max-width: 940px
    background-color: #ffffff
    box-shadow: 0 1px 0 #e0e0e0
    .title
      padding: 0 20px
      font-size: 1.6rem
      a, a:link, a:visited, a:hover, a:active
        color: $fontColor
        text-decoration: none
        cursor: pointer
    .created-at
      padding: 10px 20px
      height: 50px
      font-size: .9rem
      color: lighten($fontColor, 10%)
      .time
        margin-top: -5px
        font-size: 1rem
    .content
      padding: 0 20px
      letter-spacing: 1px
      line-height: 200%
  
  .post.unpublished
    .title a
      color: lighten($fontColor, 40%)

  .no-posts, .loading-posts
    padding: 60px 0 0 0
    text-align: center
    font-size: 32px
    color: lighten($fontColor, 20%)
</style>

<style lang="sass">
#posts-view
  .post
    .content p
      margin: 0
</style>
