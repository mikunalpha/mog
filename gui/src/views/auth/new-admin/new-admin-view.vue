<template>
  <div id="new-admin-view">
    <div class="title">
      Create Admin
    </div>

    <div class="panel">
      <ui-textbox
        class="username"
        type="text"
        icon="person"
        placeholder="Username"
        v-model="account.username">
      </ui-textbox>

      <ui-textbox
        class="password"
        type="password"
        icon="lock"
        placeholder="Password"
        v-model="account.password">
      </ui-textbox>

      <ui-textbox
        class="password"
        type="password"
        icon="lock"
        placeholder="Confirm Password"
        v-model="confirmPassword">
      </ui-textbox>

      <br/>

      <div class="actions">
        <div class="left">

        </div>
        <div class="right">
          <ui-button
            type="primary"
            color="primary"
            :loading="isSaving"
            @click="send">
            Send
          </ui-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'new-admin-view',

  data () {
    return {
      account: {
        username: '',
        password: ''
      },
      confirmPassword: '',
      isSaving: false
    }
  },

  computed: {
    ...mapGetters({
      status: 'status',
      roles: 'roles'
    })
  },

  methods: {
    ...mapActions({
      createAdmin: 'createAdmin',
      getStatus: 'getStatus'
    }),

    send () {
      if (this.account.password !== this.confirmPassword) {
        return
      }

      this.isSaving = true
      this.account.role = this.roles.Admin

      this.createAdmin({
        account: this.account,
        success: (account) => {
          this.$router.replace({name: 'Auth.Login'})
          this.getStatus({})
        },
        error: (status, e) => {
          this.isSaving = false
        }
      })
    }
  },

  mounted () {
    if (this.status.gotAt !== null && this.status.administered) {
      this.$router.replace({name: 'Blog.Posts'})
    }
  }
}
</script>

<style lang="sass">
#new-admin-view
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
    width: 360px
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
