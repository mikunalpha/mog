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
      <panel class="top-panel">
        <template slot="right">
          <ui-icon-button
            class="action-button"
            type="secondary"
            color="default"
            icon="add"
            @click="$router.push({name: 'Admin.Blog.NewPost'})">
          </ui-icon-button>
        </template>
      </panel>

      <div class="scroll-wrap">
        <posts-list
          :posts="shownPosts"
          @deletePost="deletePost">
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
      </div>
    </template>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import api from '@/api/posts'
import Panel from '@/components/admin/panel'
import PostsList from '@/components/admin/posts-list'

export default {
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
      perPagePostsNumber: 5
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
    Panel,
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
    ...mapActions({
      enableKeepAlive: 'enableKeepAlive'
    }),

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
                name: 'Admin.Blog.Post',
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
    },

    deletePost ({id, post}) {
      let self = this

      api.deletePost({
        id,
        successCallback: (post) => {
          // reload
          self.goToPage(self.currentPage)
        },
        errorCallback: ({status, error}) => {
          console.log(status)
          console.log(error)
        }
      })
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
  },

  mounted () {
    this.enableKeepAlive()
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

$scrollWrapPaddingTop: $adminLogoHeight - 10px

#posts-view
  position: relative
  height: 100%
  max-height: 100vh

  .extra-info
    padding-top: 50px
    font-size: 24px
    color: lighten($fontColor, 20%)    
    text-align: center

  .top-panel
    position: absolute
    top: 0
    left: 0
    z-index: 99
    width: 100%
    box-shadow: 0 -1px 0 darken($backgroundColor, 10%) inset

    .action-button
      margin-right: 10px

  .scroll-wrap
    padding-top: $scrollWrapPaddingTop
    height: 100vh
    max-height: 100vh
    overflow-y: auto

  .pagination
    margin-top: 20px
    margin-bottom: 20px
    text-align: center
    
    .current-page
      padding: 0 20px
      user-select: none
</style>
