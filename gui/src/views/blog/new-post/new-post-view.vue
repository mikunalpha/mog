<template>
  <div id="new-post-view">
    <div class="wrap edit">
      <input
        class="title"
        type="text"
        placeholder="Insert title here..."
        v-model="newPost.title"
        @keyup=""
        @cut=""
        @paste="">

      <quill-editor
        class="editor"
        placeholder="Insert content here..."
        :options="quillEditorOptions"
        v-model="newPost.content"
        @change="">
      </quill-editor>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import { quillEditor } from 'vue-quill-editor'

export default {
  data () {
    return {
      newPost: {
        title: '',
        content: ''
      },
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
      }
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
      getAuthInfo: 'getAuthInfo'
    })
  },

  mounted () {
    this.getAuthInfo({}) // will be deprecated
  }
}
</script>

<style lang="sass">
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
    .title
      padding: 0 20px
      width: 100%
      text-align: center
      font-size: 2rem
      font-family: Roboto
      border: none
    .editor
      margin-top: 20px
      .ql-container
        border: none
      .ql-editor
        padding: 20px
        max-height: 60vh
        min-height: 300px
        letter-spacing: 1px
        background-color: #ffffff
        font-size: 1.2em
        line-height: 200%
</style>
