<template>
  <div id="post-view">
    <template v-if="!isReady">
      Loading ...
    </template>
    <template v-else-if="errorMessage !== ''">
      {{ errorMessage }}
    </template>
    <template v-else>
      <panel class="top-panel">
        <div class="editor-toolbar" slot="left">
          <span class="ql-formats">
            <select class="ql-header">
              <option value="1">Header 1</option>
              <option value="2">Header 2</option>
              <option value="3">Header 3</option>
              <option value="4">Header 4</option>
              <option value="5">Header 5</option>
              <option value="6">Header 6</option>
            </select>
          </span>

          <span class="ql-formats">
            <button class="ql-bold"></button>
            <button class="ql-italic"></button>
            <button class="ql-underline"></button>
            <button class="ql-strike"></button>
          </span>

          <span class="ql-formats">
            <button class="ql-script" value="sub"></button>
            <button class="ql-script" value="super"></button>
          </span>

          <span class="ql-formats">
            <button class="ql-link"></button>
            <button class="ql-image"></button>
          </span>

          <span class="ql-formats">
            <button class="ql-blockquote"></button>
            <button class="ql-code-block"></button>
          </span>

          <span class="ql-formats">
            <button class="ql-clean"></button>
          </span>
        </div>

        <template slot="right">
          <switch-button class="action-button" 
            :options="switchOptions"
            @select="changePublishStatus">
          </switch-button>

          <ui-icon-button
            class="action-button"
            type="secondary"
            color="default"
            icon="save"
            :loading="updating"
            @click="updatePost">
          </ui-icon-button>
        </template>
      </panel>
      
      <div class="scroll-wrap">
        <div class="editor-wrap">
          <input
            class="title"
            type="text"
            placeholder="Insert title here..."
            v-model="post.title">

          <quill-editor
            class="editor-content"
            placeholder="Insert content here..."
            :options="quillEditorOptions"
            v-model="post.content">
          </quill-editor>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import api from '@/api/posts'
import Panel from '@/components/admin/panel'
import SwitchButton from '@/components/switch'
import { quillEditor } from 'vue-quill-editor'

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
      switchOptions: [
        {
          text: 'Unpublsihed',
          value: false,
          selected: true
        },
        {
          text: 'Publsihed',
          value: true,
          selected: false
        }
      ],
      errorMessage: '',
      updating: false,
      quillEditorOptions: {
        modules: {
          toolbar: '.editor-toolbar',
          syntax: true
        }
      }
    }
  },

  computed: {
    isReady () {
      return this.gotAt !== null
    }
  },

  components: {
    Panel,
    SwitchButton,
    quillEditor
  },

  methods: {
    ...mapActions({
      disableKeepAlive: 'disableKeepAlive'
    }),

    changePublishStatus ({index, option}) {
      this.post.published = option.value
    },

    updatePost () {
      let self = this

      if (self.post.id === '') {
        return
      } else if (self.post.title === '') {
        return
      } else if (self.post.content === '') {
        return
      }

      self.updating = true

      setTimeout(() => {
        api.updatePost({
          post: self.post,
          successCallback: () => {
            self.updating = false
            self.disableKeepAlive()
          },
          errorCallback: ({status, error}) => {
            self.updating = false
            self.disableKeepAlive()
          }
        })
      }, 500)
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
        if (self.post.published) {
          self.switchOptions[0].selected = false
          self.switchOptions[1].selected = true
        }
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

$editorWrapPaddingTop: $adminLogoHeight + 10px

#post-view
  position: relative
  height: 100%
  max-height: 100vh

  .top-panel
    position: absolute
    top: 0
    left: 0
    z-index: 99
    width: 100%
    box-shadow: 0 -1px 0 darken($backgroundColor, 10%) inset

    .editor-toolbar
      border: none

    .ql-formats
      margin-right: 8px
    
    .action-button
      margin-right: 10px
    
    .publishment-switch

  .scroll-wrap
    height: 100vh
    max-height: 100vh
    overflow-y: auto

  .editor-wrap
    margin: 0 auto
    padding: 20px 30px
    padding-top: $editorWrapPaddingTop
    width: 100%
    background-color: $backgroundColor
    @media screen and (max-width: $smallScreenWidth)
      width: 100%
      max-width: 100%
    @media screen and (min-width: $largeScreenWidth)
      width: 80%
      max-width: 920px
      padding: 30px 50px
      padding-top: $editorWrapPaddingTop

    .title
      display: block
      margin: 0 auto
      padding-bottom: 5px
      width: 100%
      font-family: 'Roboto'
      font-size: 24px
      font-weight: bold
      border: none
      box-shadow: 0 -1px 0 lighten($fontColor, 20%) inset

    .editor-content
      margin: 0 auto
      // margin-top: 20px
      width: 100%
      max-width: 100%
      letter-spacing: 1px
      line-height: 160%
    
    .ql-container
      font-size: 16px
      border: none

      .ql-editor
        padding: 0
        &::before
          margin-top: 16px
          font-style: normal

        p, ul, ol
          margin-top: 16px
          margin-bottom: 16px

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

#post-view .editor-wrap
  .ql-container
    font-size: 16px
    border: none

    .ql-editor
      padding: 0
      &::before
        margin-top: 16px
        font-style: normal

      p, ul, ol
        margin-top: 16px
        margin-bottom: 16px

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
