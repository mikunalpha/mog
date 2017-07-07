<template>
  <div id="login-view">
    <template v-if="status && status.administered">
    <div class="title">
      Mog
    </div>

    <div class="panel">
      <ui-textbox
        class="username"
        type="text"
        icon="person"
        placeholder="Username"
        v-model="credentials.username">
      </ui-textbox>

      <ui-textbox
        class="password"
        type="password"
        icon="lock"
        placeholder="Password"
        v-model="credentials.password">
      </ui-textbox>

      <br/>

      <div class="actions">
        <div class="left">

        </div>
        <div class="right">
          <ui-button
            type="primary"
            color="primary"
            :loading="isLogining"
            @click="login">
            Login
          </ui-button>
        </div>
      </div>
    </div>
    </template>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data () {
    return {
      credentials: {
        username: '',
        password: ''
      },
      isLogining: false
    }
  },

  methods: {
    ...mapActions({
      loginAccount: 'loginAccount',
      getAuthInfo: 'getAuthInfo'
    }),

    login () {
      this.isLogining = true

      this.loginAccount({
        credentials: this.credentials,
        success: (data) => {
          this.getAuthInfo({})
          this.$router.push({name: 'Blog.Posts'})
        },
        error: (status, e) => {
          this.isLogining = false
          console.log(status)
        }
      })
    }
  },

  computed: {
    ...mapGetters({
      status: 'status'
    })
  }
}
</script>

<style lang="sass">
#login-view
  padding-top: 100px
  min-height: 100vh
  background-color: #f0f0f0
  .title
    padding-bottom: 20px
    text-align: center
    font-size: 1.6rem
  .panel
    margin: 0 auto
    padding: 40px 30px 30px 30px
    width: 300px
    background-color: #ffffff
    box-shadow: 0 1px 0 #e0e0e0
    .username, .password
      input
        padding-left: 5px
        letter-spacing: 4px
    .actions
      display: flex
      > div
        flex: 1
      .right
        text-align: right
</style>
