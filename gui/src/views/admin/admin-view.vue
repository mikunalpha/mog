<template>
  <div id="admin-view">
    <template v-if="!isReady">
      <!-- Loading ... -->
    </template>
    <template v-else>
      <aside class="admin-side-panel">
        <div class="logo">
          Mog Admin
        </div>
        <navigation :items="navigationItems"></navigation>
      </aside>

      <keep-alive>
        <router-view v-if="$route.meta.keepAlive"></router-view>
      </keep-alive>
      <router-view v-if="!$route.meta.keepAlive"></router-view>
    </template>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Navigation from '@/components/admin/navigation'

export default {
  data () {
    return {
      navigationItems: [
        {
          type: 'link',
          icon: 'dashboard',
          text: 'Overview',
          routerLink: {
            name: 'Admin.Overview'
          }
        },
        {
          type: 'link',
          text: 'Accounts',
          icon: 'people',
          routerLink: {
            name: 'Admin.Accounts'
          }
        },
        {
          type: 'label',
          text: 'Blog',
          routerLink: null,
          children: [
            {
              type: 'link',
              text: 'Posts',
              icon: 'library_books',
              routerLink: {
                name: 'Admin.Blog.Posts',
                query: {
                  page: 1
                }
              }
            }
          ]
        }
      ]
    }
  },

  computed: {
    ...mapGetters({
      status: 'status',
      authInfo: 'authInfo',
      keepAlive: 'keepAlive'
    }),

    isReady () {
      return (this.status.gotAt !== null) && (this.authInfo.gotAt !== null)
    }
  },

  components: {
    Navigation
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#admin-view
  display: flex
  position: relative
  // padding-bottom: 1px
  min-height: 100vh
  background-color: darken($backgroundColor, 5%)
  overflow-x: hidden
  
  > div
    flex-grow: 1

  .admin-side-panel
    // width: 300px
    max-width: 300px
    background-color: darken($backgroundColor, 60%)

    .logo
      display: flex
      align-items: center
      padding-left: 20px
      height: $adminLogoHeight
      font-size: 24px
      color: lighten($backgroundColor, 100%)
      font-weight: bold
      user-select: none
</style>
