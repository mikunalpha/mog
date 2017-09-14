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

    <div class="excerpt" v-html="post.content">
    </div>
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
  }
}
</script>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

.posts-list-item
  margin: 20px auto
  padding: 20px
  width: 80%
  // border: 1px solid darken($backgroundColor, 10%)
  background-color: $backgroundColor

  @media screen and (max-width: $smallScreenWidth)
    width: 100%
  @media screen and (min-width: $largeScreenWidth)
    width: 920px
    padding: 30px 30px 20px 30px

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
  
  .excerpt
    letter-spacing: 1px
    line-height: 160%
  
  .ql-syntax
    margin: 5px 0
    padding: 10px
    background-color: darken($backgroundColor, 5%)
    color: $fontColor
  
  .ql-align-left
    text-align: left
  .ql-align-right
    text-align: right
  .ql-align-center
    text-align: center
  .ql-align-justify
    text-align: justify
</style>
