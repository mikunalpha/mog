import Vue from 'vue'
import Router from 'vue-router'
import Store from '@/store'

import BadRequestView from '@/views/bad-request-view'
import NotFoundView from '@/views/not-found-view'
import ServerErrorView from '@/views/server-error-view'

import AuthView from '@/views/auth/auth-view'
import AuthLoginView from '@/views/auth/login/login-view'
import AuthCreateAdminView from '@/views/auth/create-admin/create-admin-view'

// import BlogView from '@/views/blog/blog-view'
// import BlogPostsView from '@/views/blog/posts/posts-view'
// import BlogPostsActions from '@/views/blog/posts/posts-actions'
// import BlogPostView from '@/views/blog/post/post-view'
// import BlogPostActions from '@/views/blog/post/post-actions'
// import BlogNewPostView from '@/views/blog/new-post/new-post-view'
// import BlogNewPostActions from '@/views/blog/new-post/new-post-actions'

import BlogView from '@/views/blog/blog-view'
import BlogPostsView from '@/views/blog/posts/posts-view'
import BlogPostView from '@/views/blog/posts/post-view'

import AdminView from '@/views/admin/admin-view'
import AdminOverviewView from '@/views/admin/overview/overview-view'
import AdminAccountsView from '@/views/admin/accounts/accounts-view'
import AdminBlogPostsView from '@/views/admin/blog/posts/posts-view'
import AdminBlogPostView from '@/views/admin/blog/posts/post-view'
import AdminBlogNewPostView from '@/views/admin/blog/posts/new-post-view'

let routes = [
  {
    path: '*',
    redirect: to => {
      return {
        name: '404'
      }
    }
  },
  {
    name: '400',
    path: '/400',
    component: BadRequestView
  },
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
        component: AuthCreateAdminView
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
        meta: {
          keepAlive: true
        },
        component: BlogPostsView,
        props: (route) => ({
          query: route.query
        })
      },
      {
        name: 'Blog.Post',
        path: 'post/:id',
        component: BlogPostView,
        props: (route) => ({
          postId: route.params.id
        })
      }
    ]
  },
  {
    name: 'Admin',
    path: '/admin',
    component: AdminView,
    children: [
      {
        name: 'Admin.Overview',
        path: 'overview',
        component: AdminOverviewView
      },
      {
        name: 'Admin.Accounts',
        path: 'accounts',
        component: AdminAccountsView
      },
      {
        name: 'Admin.Blog.Posts',
        path: 'blog/posts',
        meta: {
          keepAlive: true
        },
        beforeEnter: (to, from, next) => {
          to.meta.keepAlive = Store.getters.keepAlive
          next()
        },
        component: AdminBlogPostsView,
        props: (route) => ({
          query: route.query
        })
      },
      {
        name: 'Admin.Blog.Post',
        path: 'blog/post/:id',
        component: AdminBlogPostView,
        props: (route) => ({
          postId: route.params.id
        })
      },
      {
        name: 'Admin.Blog.NewPost',
        path: 'blog/new/post',
        component: AdminBlogNewPostView
      }
    ]
  }
]

Vue.use(Router)

var router = new Router({
  mode: 'history',
  routes,
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
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
