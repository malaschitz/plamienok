<!--
  - Developed by Richard Malaschitz on 3/5/19 2:34 PM
  - Last modified 3/5/19 2:34 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <h2 class="headline">
      {{ person.FirstName }} {{ person.Surname }} <span class="subheading">{{ age }}</span>
    </h2>
    <v-tabs
      v-model="active"
      color="primary"
      dark
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
        Poradňa
      </v-tab>

      <v-tab-item key="0">
        <v-form v-model="valid" ref="baseform" lazy-validation>
          <v-container>
            <v-layout wrap>
              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.FirstName"
                  :rules="nameRules"
                  :counter="10"
                  label="Meno"
                  required
                />
              </v-flex>
              <v-flex xs12 md4 sm6>
                <v-text-field
                  v-model="person.Surname"
                  :rules="nameRules"
                  :counter="10"
                  label="Priezvisko"
                  required
                />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.RC" label="Rodné číslo" mask="######/####" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-checkbox v-model="person.IsHC" label="HomeCare" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-checkbox v-model="person.IsCGT" label="CGT" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-checkbox v-model="person.IsPatient" label="Klient" />
              </v-flex>

              <v-flex xs12 md3 sm6>
                <v-select
                  v-model="person.Sex"
                  :items="[{value:0,text: '-'},{value:1,text: 'muž'},{value:2,text: 'žena'}]"
                  label=""
                />
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker1"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  lazy
                  transition="scale-transition"
                  offset-y
                  full-width
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="person.DtoBirthDate"
                      label="Dátum narodenia"
                      prepend-icon="event"
                      readonly
                      v-on="on"
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker v-model="person.DtoBirthDate" @input="menuDatePicker1 = false" />
                </v-menu>
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker2"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  lazy
                  transition="scale-transition"
                  offset-y
                  full-width
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="person.DtoDeath"
                      label="Dátum úmrtia"
                      prepend-icon="event"
                      readonly
                      v-on="on"
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker v-model="person.DtoDeath" @input="menuDatePicker2 = false" />
                </v-menu>
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker3"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  lazy
                  transition="scale-transition"
                  offset-y
                  full-width
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="person.DtoPlamPrijatie"
                      label="Dátum prijatia"
                      prepend-icon="event"
                      readonly
                      v-on="on"
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker v-model="person.DtoPlamPrijatie" @input="menuDatePicker3 = false" />
                </v-menu>
              </v-flex>

              <v-flex xs12 sm6 md3>
                <v-menu
                  v-model="menuDatePicker4"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  lazy
                  transition="scale-transition"
                  offset-y
                  full-width
                  min-width="290px"
                >
                  <template v-slot:activator="{ on }">
                    <v-text-field
                      v-model="person.DtoPlamPrepustenie"
                      label="Dátum prepustenia"
                      prepend-icon="event"
                      v-on="on"
                      readonly
                      required
                      clearable
                    />
                  </template>
                  <v-date-picker v-model="person.DtoPlamPrepustenie" @input="menuDatePicker4 = false" />
                </v-menu>
              </v-flex>
            </v-layout>

            <h3>Kontaktné údaje</h3>
            <v-layout wrap>
              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.Email" :rules="emailRules" label="E-mail" />
              </v-flex>
              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.Phone" label="Telefónne číslo" />
              </v-flex>
              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.Job" label="Zamestnanie" />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.AddrStreet" label="Ulica" />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.AddrCity" label="Mesto" />
              </v-flex>

              <v-flex xs12 md4 sm6>
                <v-text-field v-model="person.AddrPSC" label="PSČ" mask="### ##" />
              </v-flex>

              <v-btn color="info" @click="saveBase">
                Uložiť
              </v-btn>

              <v-btn color="warning" @click="readData">
                Refresh
              </v-btn>
            </v-layout>
          </v-container>
        </v-form>
      </v-tab-item>

      <v-tab-item key="1">
        <v-form v-model="valid" ref="zuform" lazy-validation>
          <v-container>
            <v-layout wrap>
              <v-flex xs12 md4 sm6>
                <v-select
                  v-model="person.ZP"
                  :items="[{value:'',text: '-'},{value:'VZP',text: 'Všeobecná ZP'},{value:'dovera',text: 'Dôvera'},{value:'union',text: 'Union'}]"
                  label=""
                />
                <v-flex xs12 md8 sm6>
                  <v-text-field v-model="person.Odoslal" label="Odosielajúci lekár" />
                </v-flex>
                <v-flex xs12 md12 sm12>
                  <v-autocomplete
                    v-model="model"
                    :items="items"
                    :loading="isLoading"
                    :search-input.sync="search"
                    chips
                    clearable
                    hide-details
                    hide-selected
                    item-text="Popis"
                    item-value="ID"
                    label="Zadanie diagnózy"
                  >
                    <template v-slot:no-data>
                      <v-list-tile>
                        <v-list-tile-title>
                          Vyhľadávanie diagnózy
                        </v-list-tile-title>
                      </v-list-tile>
                    </template>
                    <template v-slot:selection="{ item, selected }">
                      <v-chip
                        :selected="selected"
                        color="blue-grey"
                        class="white--text"
                      >
                        <v-icon left>
                          mdi-coin
                        </v-icon>
                        <span v-text="item.Popis" />
                      </v-chip>
                    </template>
                    <template v-slot:item="{ item }">
                      <v-list-tile-avatar
                        color="indigo"
                        class="headline font-weight-light white--text"
                      >
                        {{ item.Skratka.charAt(0) }}
                      </v-list-tile-avatar>
                      <v-list-tile-content>
                        <v-list-tile-title v-text="item.Popis" />
                        <v-list-tile-sub-title v-text="item.Skratka" />
                      </v-list-tile-content>
                      <v-list-tile-action>
                        <v-icon>mdi-coin</v-icon>
                      </v-list-tile-action>
                    </template>
                  </v-autocomplete>
                </v-flex>

                <v-btn color="info" @click="saveZP">
                  Uložiť
                </v-btn>

                <v-btn color="warning" @click="readData">
                  Refresh
                </v-btn>
              </v-flex>
            </v-layout>
          </v-container>
        </v-form>
      </v-tab-item>

      <v-tab-item key="2">
        <v-card flat>
          <v-card-text>Návštevy</v-card-text>
        </v-card>
      </v-tab-item>

      <v-tab-item key="3">
        <v-card flat>
          <v-card-text>Telefonáty</v-card-text>
        </v-card>
      </v-tab-item>

      <v-tab-item key="4">
        <v-card flat>
          <v-card-text>Poradňa</v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>
  </v-container>
</template>

<script>
    export default {
        name: "Person",

        data: () => ({
          person: {},
          active: 0,
          valid: false,

          nameRules: [
              v => !!v || 'Meno je vyžadované',
          ],

          emailRules: [
              // eslint-disable-next-line
              v => ( /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || !v)  || 'E-mail must be valid'
          ],

            menuDatePicker1: false,
            menuDatePicker2: false,
            menuDatePicker3: false,
            menuDatePicker4: false,

            isLoading: false,
            items: [],
            model: null,
            search: null

        }),

        watch: {
            search(val) {
                console.log(val);

                this.isLoading = true

                // Lazily load input items
                fetch('/api/diagnozy?filter=' + val)
                    .then(res => res.json())
                    .then(res => {
                        this.items = res.data
                    })
                    .catch(err => {
                        console.log(err)
                    })
                    .finally(() => (this.isLoading = false))
            }
        },

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

            saveBase: function() {
                if (!this.$refs.baseform.validate()) {
                    return;
                }
                this.$axios.put("/r/api/person",this.person).then(response => {
                  if (response.status == 202) {
                      this.$store.commit("alert","Chyba: " + response.data.Error);
                  } else { //OK
                      this.person = response.data;
                  }
              }).catch(response => {
                  console.log('WRONG',response);
                  this.$store.commit("alert","Chyba: " + response);
              });
            },

            saveZP: function() {
                if (!this.$refs.zuform.validate()) {
                    return;
                }
                this.$axios.put("/r/api/person",this.person).then(response => {
                    if (response.status == 202) {
                        this.$store.commit("alert","Chyba: " + response.data.Error);
                    } else { //OK
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

        computed: {
            age: function() {
                var dob = this.person.DtoBirthDate;
                if (dob && dob.length === 10) {
                    var year = Number(dob.substr(0, 4));
                    var month = Number(dob.substr(4, 2)) - 1;
                    var day = Number(dob.substr(6, 2));
                    var today = new Date();
                    var age = today.getFullYear() - year;
                    if (today.getMonth() < month || (today.getMonth() == month && today.getDate() < day)) {
                        age--;
                    }
                    return "(" + age + "r)";
                } else {
                    return ""
                }
            },
        }

    }
</script>

<style scoped>

</style>