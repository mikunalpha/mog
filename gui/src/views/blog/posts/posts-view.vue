<template>
  <div id="posts-view">
    <template v-if="!isReady">
      <div class="extra-info">
        Loading ...
      </div>
    </template>
    <template v-else-if="posts.length === 0">
      <div class="extra-info">
        No Posts Here
      </div>
    </template>
    <template v-else>
      <posts-list
        :posts="shownPosts">
      </posts-list>

      <div class="pagination">
        <ui-icon-button
          icon="keyboard_arrow_left"
          size="large"
          type="secondary"
          :disabled="(currentPage == 1) ? true : false"
          tooltip="Previous Page"
          @click="changeToPage(currentPage - 1)">
        </ui-icon-button>
        <span class="current-page">
          {{ currentPage }}
        </span>
        <ui-icon-button
          icon="keyboard_arrow_right"
          size="large"
          type="secondary"
          :disabled="(posts.length <= perPagePostsNumber) ? true : false"
          tooltip="Next Page"
          @click="changeToPage(currentPage + 1)">
        </ui-icon-button>
      </div>
    </template>
  </div>
</template>

<script>
import api from '@/api/posts'
import PostsList from '@/components/blog/posts-list'

export default {
  name: 'posts-view',

  props: {
    query: {
      type: Object,
      default () { return {} }
    }
  },

  data () {
    return {
      gotAt: null,
      posts: [],
      perPagePostsNumber: 3
    }
  },

  computed: {
    isReady () {
      return this.gotAt !== null
    },

    currentPage () {
      let n = parseInt(this.query.page)
      if (n === 0 || !n) {
        n = 1
      }
      return n
    },

    shownPosts () {
      let shownPosts = []
      for (let i in this.posts) {
        if (i < this.perPagePostsNumber) {
          shownPosts.push(this.posts[i])
        }
      }
      return shownPosts
    }
  },

  components: {
    PostsList
  },

  watch: {
    'query.page': function (newPage, oldPage) {
      // When page change is between pages of posts.
      if (typeof newPage === 'undefined' || typeof oldPage === 'undefined') {
        return
      }
      this.goToPage(this.currentPage)
    }
  },

  methods: {
    goToPage (n) {
      if (n <= 0) {
        return
      }
      let self = this
      let options = {
        limit: self.perPagePostsNumber + 1,
        offset: (n - 1) * self.perPagePostsNumber,
        excerpt: 1
      }
      if (self.query.keyword && self.query.keyword !== '') {
        options.keyword = self.query.keyword
      }

      api.fetchPosts({
        options,
        successCallback: (posts) => {
          self.posts = posts
          for (let i in self.posts) {
            self.$set(
              self.posts[i],
              'routerLink',
              {
                name: 'Blog.Post',
                params: {
                  id: self.posts[i].id
                }
              }
            )
          }
          self.gotAt = Date.now()
          window.scrollTo(0, 0)
        },
        errorCallback: ({status, error}) => {
          self.posts = []
          self.gotAt = Date.now()
        }
      })
    },

    changeToPage (n) {
      let self = this

      if (self.currentPage === n) {
        return
      }

      // Trigger watcher of '$router.params.page'.
      self.$router.push(
        {
          query: {
            page: n,
            keyword: this.query.keyword
          }
        }
      )
    }
  },

  created () {
    let self = this

    if (!self.query.page) {
      self.$router.push(
        {
          query: {
            page: 1,
            keyword: self.query.keyword
          }
        },
        () => {
          self.goToPage(1)
        }
      )
      return
    }
    self.goToPage(self.currentPage)
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#posts-view
  .extra-info
    padding-top: 50px
    font-size: 24px
    color: lighten($fontColor, 20%)    
    text-align: center

  .pagination
    margin-top: 50px
    margin-bottom: 30px
    text-align: center
    
    .current-page
      padding: 0 20px
      user-select: none
</style>
