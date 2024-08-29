<template>
  <form @submit.prevent="handleSubmit">
    <div>
      <label for="name">Name:</label>
      <input v-model="instance.name" id="name" type="text" required />
    </div>
    <div>
      <label for="port">Port:</label>
      <input v-model="instance.port" id="port" type="number" required />
    </div>
    <button type="submit">Save</button>
    <p v-if="loading">Processing...</p>
  </form>
</template>

<script>
export default {
  props: {
    instance: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      loading: false
    }
  },
  methods: {
    handleSubmit() {
      this.loading = true;
      this.$emit('submit', { ...this.instance });
    }
  }
}
</script>
