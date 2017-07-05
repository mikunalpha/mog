import Vue from 'vue'
import Router from 'vue-router'

import NotFoundView from '@/views/not-found-view'

import BlogView from '@/views/blog/blog-view'
import BlogPostsView from '@/views/blog/posts/posts-view'
import BlogPostsActions from '@/views/blog/posts/posts-actions'
import BlogPostView from '@/views/blog/post/post-view'
import BlogPostActions from '@/views/blog/post/post-actions'
import BlogNewPostView from '@/views/blog/new-post/new-post-view'
import BlogNewPostActions from '@/views/blog/new-post/new-post-actions'

// import AdminView from '@/views/admin/admin-view'
// import AdminPostsView from '@/views/admin/posts/posts-view'
// import AdminProjectsView from '@/views/admin/projects/projects-view'

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
    // {
    //   name: 'Admin',
    //   path: '/admin',
    //   component: AdminView,
    //   children: [
    //     {
    //       name: 'Admin.Posts',
    //       path: 'posts',
    //       component: AdminPostsView,
    //       beforeEnter: (to, from, next) => {
    //         to.meta.title = 'Posts'
    //         next()
    //       }
    //     }
    //   ]
    // }
  ]
})

export default router
