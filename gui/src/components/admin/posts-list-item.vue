<template>
  <div class="posts-list-item">
    <router-link v-if="post.routerLink"
      class="title"
      :class="{unpublished: !post.published}"
      tag="a"
      :to="post.routerLink">
      {{ post.title }}
    </router-link>
    <div v-else>
      {{ post.title }}
    </div>

    <div class="created-at">
      <ui-icon class="time" icon="date_range"></ui-icon>
      {{ post.created_at | toDatetime }}
    </div>

    <ui-icon-button
      class="delete-button"
      type="secondary"
      color="default"
      icon="close"
      @click="deletePost">
    </ui-icon-button>
  </div>
</template>

<script>
import hljs from 'highlight.js'
window.hljs = hljs

export default {
  props: {
    post: {
      type: Object,
      default () { return {} }
    }
  },

  methods: {
    deletePost () {
      this.$emit('deletePost', {
        id: this.post.id,
        post: this.post
      })
    }
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

.posts-list-item
  position: relative
  margin: 20px auto
  padding: 20px
  width: 80%
  // border: 1px solid darken($backgroundColor, 10%)
  background-color: $backgroundColor

  @media screen and (max-width: $smallScreenWidth)
    width: 100%
  @media screen and (min-width: $largeScreenWidth)
    width: 100%
    max-width: 920px
    padding: 30px 30px 30px 30px

  .title
    font-size: 24px
    font-weight: bold
    &, &:link, &:visited, &:hover, &:active
      color: lighten($fontColor, 5%)
      text-decoration: none
      cursor: pointer
  .title.unpublished
    color: lighten($fontColor, 40%)
  
  .created-at
    padding: 15px 0 5px 0
    font-size: .9rem
    color: lighten($fontColor, 10%)
    .time
      margin-top: -4px
      font-size: 1rem

  .delete-button
    display: none
    position: absolute
    top: 5px
    right: 5px

    .ui-icon-button__icon
      color: lighten($fontColor, 10%)

.posts-list-item:hover

  .delete-button
    display: flex
</style>
