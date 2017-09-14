<template>
  <div id="post-view">
    <template v-if="!isReady">
      Loading ...
    </template>
    <template v-else-if="errorMessage !== ''">
      {{ errorMessage }}
    </template>
    <template v-else>
      <div
        class="title"
        :class="{unpublished: !post.published}">
        {{ post.title }}
      </div>

      <div class="created-at">
        <ui-icon class="time" icon="date_range"></ui-icon>
        {{ post.created_at | toDatetime }}
      </div>

      <div class="content" v-html="post.content">
      </div>
    </template>
  </div>
</template>

<script>
import api from '@/api/posts'

export default {
  props: {
    postId: {
      type: String,
      default () { return '' }
    }
  },

  data () {
    return {
      gotAt: null,
      post: {
        id: null,
        title: null,
        content: null,
        published: null,
        updated_at: null,
        created_at: null
      },
      errorMessage: ''
    }
  },

  computed: {
    isReady () {
      return this.gotAt !== null
    }
  },

  created () {
    let self = this

    if (self.postId === '') {
      self.errorMessage = 'This post is not exist.'
      self.gotAt = Date.now()
      return
    }

    api.fetchPost({
      id: self.postId,
      successCallback: (post) => {
        self.post = post
        self.gotAt = Date.now()
        window.scrollTo(0, 0)
      },
      errorCallback: ({status, error}) => {
        if (status === 404) {
          self.errorMessage = 'This post is not exist.'
        } else {
          self.$router.push({name: '' + status})
          return
        }
        self.gotAt = Date.now()
      }
    })
  }
}
</script>

<style lang="sass" scoped>
@import '~@/assets/sass/variables.sass'

#post-view
  margin: 20px auto
  padding: 20px
  width: 80%
  // border: 1px solid darken($backgroundColor, 10%)
  background-color: $backgroundColor
  @media screen and (max-width: $smallScreenWidth)
    width: 100%
  @media screen and (min-width: $largeScreenWidth)
    width: 920px
    padding: 30px

  .title
    font-size: 24px
    font-weight: bold
  
  .title.unpublished
    color: lighten($fontColor, 40%)
  
  .created-at
    padding: 15px 0 5px 0
    font-size: .9rem
    color: lighten($fontColor, 10%)

    .time
      margin-top: -4px
      font-size: 1rem

  .content
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

<style lang="sass">
@import '~@/assets/sass/variables.sass'

#post-view .content
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
