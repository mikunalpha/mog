<template>
  <div class="panel">
    <div class="left">
      <slot name="left"></slot>
    </div>
    <div class="center">
      <slot name="center"></slot>
    </div>
    <div class="right">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
}
</script>

<style lang="sass">
@import '~@/assets/sass/variables.sass'

.panel
  display: flex
  height: $adminLogoHeight - 10px
  background-color: $backgroundColor

  > div.left, > div.right
    flex-grow: 1

  > div
    display: flex
    align-items: center

  > div.right
    justify-content: right
</style>
