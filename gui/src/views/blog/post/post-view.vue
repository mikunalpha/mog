<template>
  <div id="post-view">
    <template v-if="mode === 'edit' && authInfo.role === roles.Admin">
      <div class="wrap edit">
        <div class="published">
          <ui-switch
            label="Published"
            v-model="editedPost.published">
          </ui-switch>
        </div>

        <input
          class="title"
          type="text"
          placeholder="Insert title here..."
          v-model="editedPost.title"
          @keyup="editedPostChange"
          @cut="editedPostChange"
          @paste="editedPostChange">

        <quill-editor
          class="editor"
          placeholder="Insert content here..."
          :options="quillEditorOptions"
          v-model="editedPost.content"
          @change="editedPostChange">
        </quill-editor>
      </div>
    </template>
    <template v-else>
      <div class="wrap view">
        <div class="title">
          {{ post.title }}
        </div>

        <div class="created-at">
          <ui-icon class="time" icon="date_range"></ui-icon>
          {{ post.created_at | toDatetime }}
        </div>

        <div class="content" v-html="post.content">
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
          // toolbar: '#editor-toolbar'
          toolbar: [
            [{'header': [1, 2, 3, false]}],
            [{'size': ['small', false, 'large', 'huge']}],
            ['bold', 'italic', 'underline', 'strike'],
            ['link', 'image'],
            ['blockquote', 'code-block'],
            [{'list': 'ordered'}, {'list': 'bullet'}],
            ['clean']
          ]
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
      clearPost: 'clearPost'
    }),

    // communicate with actions button
    in ({cmd, data}) {
      if (cmd === 'changeMode') {
        this.mode = data
        this.editedPost = Object.assign({}, this.post)
      } else if (cmd === 'savePost') {
        this.editedPostChange()
      }
    },

    editedPostChange () {
      let t = this

      if (t.savePostTimeoutHolder !== null) {
        clearTimeout(t.savePostTimeoutHolder)
      }

      t.savePostTimeoutHolder = setTimeout(() => {
        t.$emit('channel', {cmd: 'savePost', data: t.editedPost})
        t.savePostTimeoutHolder = null
      }, 480)
    }
  },

  mounted () {
    let options = {
      id: 'abcd'
    }
    this.getPost({options})
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
      padding: 0 20px 20px 20px
      min-height: 300px
      letter-spacing: 1px
      background-color: #ffffff
      line-height: 200%
      font-size: 16px

  .wrap.edit
    position: relative
    .published
      position: absolute
      top: 10px
      right: 25px
      // width: 120px
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
</style>
