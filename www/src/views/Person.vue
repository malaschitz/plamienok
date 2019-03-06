<!--
  - Developed by Richard Malaschitz on 3/5/19 2:34 PM
  - Last modified 3/5/19 2:34 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <h1>Detail osoby {{ $route.params.id }}</h1>
</template>

<script>
    export default {
        name: "Person",

        data: () => ({
          person: {
              ID: '',
          },
        }),

        methods: {
            readData: function() {
                this.$axios.get('/r/api/person/' + this.$route.params.id).then(response => {
                    console.log('OK',response);
                    if (response.status == 202) {
                        this.$store.commit("alert","Chyba: " + response.data.Error);
                    } else {
                        this.person = response.data;
                    }
                }).catch(response => {
                    console.log('WRONG',response);
                    this.$store.commit("alert","Chyba: " + response);
                });
            },
        },

        mounted: function() {
            console.log("mounted persons");
            this.readData();
        },

    }
</script>

<style scoped>

</style>