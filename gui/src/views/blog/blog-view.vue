<template>
  <div id="blog-view">
    <template v-if="!isReady">
      <!-- Loading ... -->
    </template>
    <template v-else>
      <header class="blog-view-header">
        <router-link 
          class="logo"
          tag="a"
          :to="{name: 'Blog.Posts', query: {page: 1}}">
          Mog
        </router-link>
      </header>

      <keep-alive>
          <router-view v-if="$route.meta.keepAlive"></router-view>
      </keep-alive>
      <router-view v-if="!$route.meta.keepAlive"></router-view>
    </template>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'blog-view',

  computed: {
    ...mapGetters({
      status: 'status',
      authInfo: 'authInfo'
    }),

    isReady () {
      return (this.status.gotAt !== null) && (this.authInfo.gotAt !== null)
    }
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#blog-view
  position: relative
  padding-bottom: 1px
  min-height: 100vh
  background-color: darken($backgroundColor, 5%)

  .blog-view-header
    padding: 20px 0
    
  .blog-view-header .logo
    display: block
    margin: 0 auto
    width: 100px
    font-size: 28px
    font-weight: bold
    text-align: center
    &, &:link, &:visited, &:hover, &:active
      color: $fontColor
      text-decoration: none
      cursor: pointer
</style>
