<template>
  <div id="post-view">
    <template v-if="post.gotAt === null">
      <div class="loading-post">
        Loading...
      </div>
    </template>
    <template v-else-if="post.gotAt !== null && !post.data.id">
      <div class="no-post">
        Not Found
      </div>
    </template>
    <template v-else-if="mode === 'edit' && authInfo.role === roles.Admin">
      <div class="wrap edit">
        <div class="published">
          <ui-switch
            label="Published"
            v-model="editedPost.published"
            @change="">
          </ui-switch>
        </div>

        <input
          class="title"
          type="text"
          placeholder="Insert title here..."
          v-model="editedPost.title"
          @keyup=""
          @cut=""
          @paste="">

        <quill-editor
          ref="quillEditor"
          class="editor"
          placeholder="Insert content here..."
          :options="quillEditorOptions"
          v-model="editedPost.content"
          @ready="onEditorReady($event)">
        </quill-editor>
      </div>
    </template>
    <template v-else>
      <div class="wrap view" :class="{unpublished: !post.data.published}">
        <div class="title">
          {{ post.data.title }}
        </div>

        <div class="created-at">
          <ui-icon class="time" icon="date_range"></ui-icon>
          {{ post.data.created_at | toDatetime }}
        </div>

        <div class="content" v-html="post.data.content">
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import { quillEditor } from 'vue-quill-editor'

export default {
  data () {
    return {
      mode: 'view',
      quillEditorOptions: {
        modules: {
          toolbar: [
            [{'header': [1, 2, 3, false]}],
            ['bold', 'italic', 'underline', 'strike'],
            ['link', 'image'],
            [{'align': []}],
            ['blockquote', 'code-block'],
            ['clean']
          ],
          syntax: true
        }
      },
      editedPost: {
        title: '',
        content: '',
        published: false
      },
      savePostTimeoutHolder: null
    }
  },

  computed: {
    ...mapGetters({
      authInfo: 'authInfo',
      roles: 'roles',
      post: 'post'
    })
  },

  components: {
    quillEditor
  },

  methods: {
    ...mapActions({
      getPost: 'getPost',
      refreshPost: 'refreshPost',
      clearPost: 'clearPost'
    }),

    // communicate with actions button
    in ({cmd, data}) {
      if (cmd === 'changeMode') {
        this.mode = data
        if (this.mode === 'edit') {
          this.editedPost = Object.assign({}, this.post.data)
        }
      } else if (cmd === 'updatePost') {
        this.editedPostChange()
      } else if (cmd === 'refreshPost') {
        this.refreshPostView()
      }
    },

    editedPostChange () {
      let t = this

      if (t.savePostTimeoutHolder !== null) {
        clearTimeout(t.savePostTimeoutHolder)
      }

      t.$refs.quillEditor.quill.emitter.emit('text-change')

      t.savePostTimeoutHolder = setTimeout(() => {
        t.$emit('channel', {cmd: 'updatePost', data: t.editedPost})
        t.savePostTimeoutHolder = null
      }, 480)
    },

    refreshPostView () {
      this.refreshPost({post: this.editedPost})
    },

    // workaround to remove extra tag p before code-block
    onEditorReady (editor) {
      let qlEditor = null
      for (let i = 0; i < editor.container.childNodes.length; i++) {
        if (editor.container.childNodes[i].className === 'ql-editor') {
          qlEditor = editor.container.childNodes[i]
          break
        }
      }
      if (qlEditor === null) return
      // console.log(qlEditor)

      let nodeNames = ['H1', 'H2', 'H3', 'H4', 'H5', 'H6', 'BLOCKQUOTE', 'OL', 'UL']

      for (let i = 0; i < qlEditor.childNodes.length; i++) {
        if (!qlEditor.childNodes[i].nextSibling) {
          continue
        }
        if (qlEditor.childNodes[i].nextSibling.className === 'ql-syntax' && qlEditor.childNodes[i].innerHTML === '<br>') {
          qlEditor.removeChild(qlEditor.childNodes[i])
        } else if (nodeNames.indexOf(qlEditor.childNodes[i].nextSibling.nodeName) >= 0 && qlEditor.childNodes[i].innerHTML === '<br>') {
          qlEditor.removeChild(qlEditor.childNodes[i])
        }
      }
    }
  },

  created () {
    this.getPost({
      id: this.$route.params.id,
      success: (post) => {
        // console.log(post.content)
      }
    })
  },

  mounted () {
    // scroll to top
    window.scrollTo(0, 0)
  },

  beforeDestroy () {
    this.clearPost({})
  }
}
</script>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

#post-view
  position: relative
  padding-top: 20px
  padding-bottom: 120px
  min-height: 100vh
  background-color: #f0f0f0
  .wrap
    margin: 0 auto
    padding: 40px 20px 40px 20px
    max-width: 940px
    background-color: #ffffff
  .wrap.view
    .title
      padding: 0 20px
      text-align: center
      font-size: 2rem
    .created-at
      padding: 10px 0
      height: 50px
      text-align: center
      font-size: .9rem
      color: lighten($fontColor, 10%)
      .time
        margin-top: -3px
        font-size: 1rem
    .content
      margin: 0 auto
      padding: 16px 20px 20px 20px
      min-height: 300px
      letter-spacing: 1px
      background-color: #ffffff
      line-height: 200%
      font-size: 16px
      h1, h2, h3, h4, h5, h6
        margin: 0
      p
        margin: 0
      blockquote
        margin: 5px 0
        padding-left: 16px
        border-left: 4px solid #ccc
      .ql-syntax
        margin: 5px 0
        padding: 5px 10px
        background-color: #f0f0f0
        color: $fontColor
    .view
      margin-top: 8px

  .wrap.view.unpublished
    .title
      color: lighten($fontColor, 40%)

  .wrap.edit
    position: relative
    .published
      position: absolute
      top: 10px
      right: 25px
      .ui-switch__thumb
        z-index: 1
      .ui-switch__track
        z-index: 0
    .title
      padding: 0 20px
      width: 100%
      text-align: center
      font-size: 2rem
      font-family: Roboto
      border: none
    .editor
      margin-top: 8px
      .ql-container
        border: none
      .ql-editor
        padding: 16px 20px 20px 20px
        max-height: 60vh
        min-height: 300px
        letter-spacing: 1px
        background-color: #ffffff
        font-size: 16px
        line-height: 200%
        .ql-syntax
          background-color: #f0f0f0
          color: $fontColor

  .no-post, .loading-post
    padding: 60px 0 0 0
    text-align: center
    font-size: 32px
    color: lighten($fontColor, 20%)
</style>
