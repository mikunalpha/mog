import Vue from 'vue'
import Router from 'vue-router'
import Store from '@/store'

import NotFoundView from '@/views/not-found-view'

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
      name: 'Auth.Login',
      path: '/auth/login',
      component: AuthLoginView
    },
    {
      name: 'Auth.NewAdmin',
      path: '/auth/new/admin',
      component: AuthNewAdminView
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

router.beforeEach((to, from, next) => {
  Store.dispatch('getAuthInfo', {
    success: (info) => {
      console.log(info)
      next()
    },
    error: () => {
      console.log('error')
      next()
    }
  })
})

export default router
