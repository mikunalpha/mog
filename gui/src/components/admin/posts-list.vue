<template>
  <div class="posts-list">
    <posts-list-item v-for="post in posts"
      :post="post"
      :key="post.id"
      @deletePost="deletePost">
    </posts-list-item>
  </div>
</template>

<script>
import PostsListItem from '@/components/admin/posts-list-item'

export default {
  props: {
    posts: {
      type: Array,
      default () { return [] }
    }
  },

  components: {
    PostsListItem
  },

  methods: {
    deletePost ({id, post}) {
      this.$emit('deletePost', {id, post})
    }
  }
}
</script>

<style lang="sass" scoped>
.posts-list
</style>
