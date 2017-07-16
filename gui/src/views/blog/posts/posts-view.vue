<template>
  <div id="posts-view">
    <template v-if="posts.gotAt !== null && posts.data.length > 0">
      <div
        class="keyword"
        v-if="$route.query && $route.query.keyword && $route.query.keyword.trim() !== ''">
        Result of "<b>{{ $route.query.keyword.trim() }}</b>"
      </div>

      <div 
        class="post" 
        :class="{unpublished: !p.published}" 
        v-for="(p, i) in posts.data"
        v-if="i < perPagePostsNumber">
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

      <div class="pagination">
        <ui-icon-button
          icon="keyboard_arrow_left"
          size="large"
          type="secondary"
          :disabled="(currentPage == 1) ? true : false"
          tooltip="Previous Page"
          @oclick="$router.push({query: {page: currentPage - 1, keyword: $route.query.keyword}})"
          @click="changeToPage(currentPage - 1)">
        </ui-icon-button>
        <span class="current-page">
          {{ currentPage }}
        </span>
        <ui-icon-button
          icon="keyboard_arrow_right"
          size="large"
          type="secondary"
          :disabled="(posts.data.length <= perPagePostsNumber) ? true : false"
          tooltip="Next Page"
          @oclick="$router.push({query: {page: currentPage + 1, keyword: $route.query.keyword}})"
          @click="changeToPage(currentPage + 1)">
        </ui-icon-button>
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
      perPagePostsNumber: 5
    }
  },

  computed: {
    ...mapGetters({
      roles: 'roles',
      posts: 'posts',
      scroll: 'scroll'
    }),

    currentPage () {
      let n = parseInt(this.$route.query.page)
      if (n === 0 || !n) {
        n = 1
      }
      return n
    }
  },

  watch: {
    '$route.query.page': function () {
      this.goToPage(this.currentPage)
    },
    '$route.query.keyword': function () {
      this.goToPage(this.currentPage)
    }
  },

  methods: {
    ...mapActions({
      getPosts: 'getPosts',
      saveScroll: 'saveScroll'
    }),

    // communicate with actions button
    in ({cmd, data}) {
      if (cmd === 'searchPosts') {
        this.saveScroll({x: 0, y: 0})
        this.$router.push({
          query: {
            page: this.currentPage,
            keyword: data
          }
        })
      }
    },

    goToPage (n) {
      if (n <= 0) {
        return
      }
      let t = this
      let options = {
        limit: t.perPagePostsNumber + 1,
        offset: (n - 1) * t.perPagePostsNumber,
        excerpt: 1
      }
      if (t.$route.query.keyword && t.$route.query.keyword !== '') {
        options.keyword = t.$route.query.keyword
      }

      t.getPosts({
        options: options,
        success: () => {
          // console.log('scrollTo', this.scroll.x, this.scroll.y)
          window.scrollTo(this.scroll.x, this.scroll.y)
        }
      })
    },

    changeToPage (n) {
      this.saveScroll({x: 0, y: 0})
      this.$router.push({query: {page: n, keyword: this.$route.query.keyword}})
    }
  },

  mounted () {
    // Go to page
    this.goToPage(this.currentPage)
  },

  beforeRouteLeave (to, from, next) {
    let x = (window.pageXOffset !== undefined) ? window.pageXOffset : (document.documentElement || document.body.parentNode || document.body).scrollLeft
    let y = (window.pageYOffset !== undefined) ? window.pageYOffset : (document.documentElement || document.body.parentNode || document.body).scrollTop
    console.log('saveScroll', x, y)
    this.saveScroll({x, y})
    next()
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
  .keyword
    margin: 0 auto
    padding: 0 10px 20px 10px
    max-width: 940px
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
  .pagination
    margin-top: 50px
    text-align: center
  .no-posts, .loading-posts
    padding: 60px 0 0 0
    text-align: center
    font-size: 32px
    color: lighten($fontColor, 20%)
</style>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

#posts-view
  .post
    .content p
      margin: 0
    .ql-syntax
        margin: 5px 0
        padding: 5px 10px
        background-color: #f0f0f0
        color: $fontColor
</style>
