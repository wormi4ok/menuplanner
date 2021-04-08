<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">
        {{ name }}
      </p>
    </header>
<!--    <div class="card-image">-->
<!--      <figure class="image">-->
<!--        <img src="../assets/carbonara.jpg" alt="Placeholder image">-->
<!--      </figure>-->
<!--    </div>-->
    <div class="card-content">
      <section>
        <b-progress :max="141" show-value>
          <template #bar>
            <b-progress-bar type="is-success" :value="35" show-value>P</b-progress-bar>
            <b-progress-bar type="is-info" :value="38">F</b-progress-bar>
            <b-progress-bar type="is-warning" :value="68" show-value>C</b-progress-bar>
          </template>
        </b-progress>
        <b-progress :value="calories" :max="1700" show-value>Calories</b-progress>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Recipe',
  props: {
    id: Number,
    name: String,
    protein: Number,
    fat: Number,
    carbs: Number,
    calories: Number,
  },
};
</script>

<style scoped>

</style>
