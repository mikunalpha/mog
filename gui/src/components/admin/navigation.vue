<template>
  <nav class="navigation">
    <template v-for="item in items">
      <div v-if="item.type === 'label'"
        class="navigation-item label">
        {{ item.text }}
      </div>
      <router-link v-else
        class="navigation-item"
        tag="a"
        active-class="active"
        :to="item.routerLink">
        <ui-icon v-if="item.icon"
          :icon="item.icon"></ui-icon>&nbsp;
        {{ item.text }}
      </router-link>

      <template v-if="item.children !== null">
        <template v-for="itemChild in item.children">
          <div v-if="itemChild.type === 'label'"
            class="navigation-item child">
            {{ itemChild.text }}
          </div>
          <router-link v-else
            class="navigation-item child"
            tag="a"
            active-class="active"
            :to="itemChild.routerLink">
            <ui-icon v-if="itemChild.icon"
              :icon="itemChild.icon"></ui-icon>&nbsp;
            {{ itemChild.text }}
          </router-link>
        </template>
        
      </template>
    </template>
  </nav>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      default () { return [] }
    }
  }
}
</script>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

.navigation
  width: 240px

  .navigation-item
    display: flex
    align-items: center
    padding-left: 20px
    height: $adminLogoHeight
    box-shadow: 0 -1px 0 darken($backgroundColor, 30%) inset
    color: darken($backgroundColor, 20%)
    font-size: 18px
    letter-spacing: 1px
    user-select: none
    &, &:link, &:visited, &:hover, &:active
      text-decoration: none
      cursor: pointer
    &.active
      color: #ffffff
      background-color: rgba(255,255,255,.225)
      box-shadow: none

  .navigation-item.label
    height: $adminLogoHeight - 10px
    font-size: 14px
    background-color: rgba(0,0,0,.125)
    cursor: default
  
  .navigation-item.child
    padding-left: 30px
</style>
