<template>
  <v-container>
    <h1>Date Loader</h1>
    <v-row justify="center" >
      <p>Pick the date you want to load:</p>
    </v-row>
    <v-row justify="center">
      <v-date-picker
        v-model="date"
        :max="getToday()"
      ></v-date-picker>
    </v-row>
    <v-row justify="center" >
      <v-btn
        outlined
        :disabled="date === null"
        v-on:click="loadDate()"
        :loading="loading">
        Load Date
      </v-btn>
    </v-row>
    <v-row justify="center">
      {{message}}
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Loaddate',
  data () {
    return {
      date: null,
      message: '',
      loading: false
    }
  },
  methods: {
    getToday: function () {
      const today = new Date()
      return today.toISOString().slice(0, 10)
    },
    loadDate: function () {
      if (this.date === null) {
        return
      }
      const timestamp = Date.parse(this.date) / 1000
      let success = true
      this.loading = true
      this.message = ''
      axios
        .post('http://localhost:5000/api/load?date=' + timestamp)
        .then(response => {
          success = response.data.success
          if (success) {
            this.message = 'The date was loaded successfully.'
          } else {
            this.message = 'The selected date was already loaded.'
          }
        })
        .catch(error => {
          console.log(error)
          this.message = 'An error ocurred while loading the data.'
        })
        .finally(() => {
          this.loading = false
        })
    }
  }

}
</script>
