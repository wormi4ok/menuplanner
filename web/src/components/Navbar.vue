<template>
  <b-navbar>
    <template #brand>
      <b-navbar-item>
        <img src="https://flaningo.ru/wp-content/uploads/2019/03/cropped-99-1.jpg" alt="MenuPlanner logo">
      </b-navbar-item>
    </template>
    <template #start>
      <b-navbar-item href="#">
        <div class="buttons">
          <a class="button is-primary">
            <strong>Fill gaps</strong>
          </a>
        </div>
      </b-navbar-item>
      <b-navbar-item href="#">
        Add recipe
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
export default {
  name: 'Navbar',
};
</script>

<style scoped>

</style>
