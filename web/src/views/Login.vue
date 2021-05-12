<template>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-5-desktop is-4-widescreen">
            <AuthForm />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import AuthForm from '@/components/AuthForm.vue';

export default {
  name: 'Login',
  components: { AuthForm },
};
</script>

<style scoped>

</style>
