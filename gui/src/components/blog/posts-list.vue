<template>
  <div class="posts-list">
    <posts-list-item v-for="post in posts"
      :post="post"
      :key="post.id">
    </posts-list-item>
  </div>
</template>

<script>
import PostsListItem from '@/components/blog/posts-list-item'

export default {
  props: {
    posts: {
      type: Array,
      default () { return [] }
    }
  },

  components: {
    PostsListItem
  }
}
</script>

<style lang="sass" scoped>
.posts-list
</style>
