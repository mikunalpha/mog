<template>
  <div id="new-post-view">
    <template v-if="authInfo.role === roles.Admin">
      <div class="wrap edit">
        <div class="published">
          <ui-switch
            label="Published"
            v-model="newPost.published">
          </ui-switch>
        </div>

        <input
          class="title"
          type="text"
          placeholder="Insert title here..."
          v-model="newPost.title"
          @keyup=""
          @cut=""
          @paste="">

        <quill-editor
          ref="quillEditor"
          class="editor"
          placeholder="Insert content here..."
          :options="quillEditorOptions"
          v-model="newPost.content"
          @change="">
        </quill-editor>
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
      quillEditorOptions: {
        modules: {
          // toolbar: '#editor-toolbar'
          toolbar: [
            [{'header': [1, 2, 3, false]}],
            ['bold', 'italic', 'underline', 'strike'],
            ['link', 'image'],
            ['blockquote', 'code-block'],
            ['clean']
          ]
        }
      },
      newPost: {
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
      roles: 'roles'
    })
  },

  components: {
    quillEditor
  },

  methods: {
    ...mapActions({
    }),

    // communicate with actions button
    in ({cmd, data}) {
      if (cmd === 'savePost') {
        this.newPostChange()
      }
    },

    newPostChange () {
      let t = this

      if (this.newPost.title.trim().length === 0) {
        // this.errorMessage = 'Username was required.'
        return
      } else {
        this.newPost.title = this.newPost.title.trim()
      }

      if (t.savePostTimeoutHolder !== null) {
        clearTimeout(t.savePostTimeoutHolder)
      }

      t.$refs.quillEditor.quill.emitter.emit('text-change')

      t.savePostTimeoutHolder = setTimeout(() => {
        t.$emit('channel', {cmd: 'savePost', data: t.newPost})
        t.savePostTimeoutHolder = null
      }, 120)
    }
  },

  mounted () {
    //
  }
}
</script>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

#new-post-view
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
</style>
