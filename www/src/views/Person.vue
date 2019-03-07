<!--
  - Developed by Richard Malaschitz on 3/5/19 2:34 PM
  - Last modified 3/5/19 2:34 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <v-tabs
      v-model="active"
      color="blue"
      dark
      slider-color="yellow"
    >
      <v-tab key="0" ripple>
        Základné údaje
      </v-tab>
      <v-tab key="1" ripple>
        Zdravotné údaje
      </v-tab>
      <v-tab key="2" ripple>
        Návštevy
      </v-tab>
      <v-tab key="3" ripple>
        Telefonáty
      </v-tab>
      <v-tab key="4" ripple>
        Stretnutia
      </v-tab>

      <v-tab-item key="0">
        <v-card flat>
          <v-card-text>Základné údaje</v-card-text>
        </v-card>
      </v-tab-item>

      <v-tab-item key="1">
        <v-card flat>
          <v-card-text>údaje</v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>
  </v-container>
</template>

<script>
    export default {
        name: "Person",

        data: () => ({
          person: {
              ID: '',
          },
          active: 0,
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