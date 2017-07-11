import Vue from 'vue'
import Router from 'vue-router'
import Store from '@/store'

import NotFoundView from '@/views/not-found-view'
import ServerErrorView from '@/views/server-error-view'

import AuthView from '@/views/auth/auth-view'
import AuthLoginView from '@/views/auth/login/login-view'
import AuthNewAdminView from '@/views/auth/new-admin/new-admin-view'

import BlogView from '@/views/blog/blog-view'
import BlogPostsView from '@/views/blog/posts/posts-view'
import BlogPostsActions from '@/views/blog/posts/posts-actions'
import BlogPostView from '@/views/blog/post/post-view'
import BlogPostActions from '@/views/blog/post/post-actions'
import BlogNewPostView from '@/views/blog/new-post/new-post-view'
import BlogNewPostActions from '@/views/blog/new-post/new-post-actions'

Vue.use(Router)

var router = new Router({
  mode: 'history',
  routes: [
    {
      name: '404',
      path: '/404',
      component: NotFoundView
    },
    {
      name: '500',
      path: '/500',
      component: ServerErrorView
    },
    {
      name: 'Auth',
      path: '/auth',
      component: AuthView,
      children: [
        {
          name: 'Auth.Login',
          path: 'login',
          component: AuthLoginView
        },
        {
          name: 'Auth.NewAdmin',
          path: 'new/admin',
          component: AuthNewAdminView
        }
      ]
    },
    {
      name: 'Blog',
      path: '/blog',
      component: BlogView,
      children: [
        {
          name: 'Blog.Posts',
          path: 'posts',
          components: {
            default: BlogPostsView,
            actions: BlogPostsActions
          },
          beforeEnter: (to, from, next) => {
            to.meta.title = 'Posts'
            dispatchEvent(new Event('load'))
            next()
          }
        },
        {
          name: 'Blog.Post',
          path: 'post/:id',
          components: {
            default: BlogPostView,
            actions: BlogPostActions
          },
          beforeEnter: (to, from, next) => {
            to.meta.title = 'Post'
            next()
          }
        },
        {
          name: 'Blog.NewPost',
          path: 'new/post',
          components: {
            default: BlogNewPostView,
            actions: BlogNewPostActions
          },
          beforeEnter: (to, from, next) => {
            to.meta.title = 'New Post'
            next()
          }
        }
      ]
    }
  ]
})

// let checkAdministerdRoutePathPrefixes = [
//   '/auth/'
// ]

let administerOnlyRoutePathPrefixes = [
  '/blog/new'
]

router.beforeEach((to, from, next) => {
  // Avoid non-admin touch some page.
  for (let i in administerOnlyRoutePathPrefixes) {
    if (to.path.startsWith(administerOnlyRoutePathPrefixes[i])) {
      // console.log(Store.getters.authInfo.role)
      if (Store.getters.authInfo.role !== Store.getters.roles.Admin) {
        next('/blog/posts')
        return
      }
    }
  }

  // Get and check status
  // if (to.path !== '/auth/new/admin') {
  //   for (let i in checkAdministerdRoutePathPrefixes) {
  //     if (to.path.startsWith(checkAdministerdRoutePathPrefixes[i])) {
  //       Store.dispatch('getStatus', {
  //         success: (status) => {
  //           console.log('= =')
  //           if (status.administered === true) {
  //             next()
  //             return
  //           }
  //           next('/auth/new/admin')
  //         }
  //       })
  //     }
  //   }
  // }

  // Store.dispatch('getAuthInfo', {
  //   success: (info) => {
  //     console.log(info)
  //     next()
  //   },
  //   error: (status, e) => {
  //     console.log(status)
  //     next()
  //   }
  // })

  next()
})

export default router
