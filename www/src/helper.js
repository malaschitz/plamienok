// define a mixin object

let myMixin = {
  created: function() {},

  methods: {
    days_between: function(date1, date2) {
      // The number of milliseconds in one day
      var ONE_DAY = 1000 * 60 * 60 * 24;

      // Convert both dates to milliseconds
      var date1_ms = date1.getTime();
      var date2_ms = date2.getTime();

      // Calculate the difference in milliseconds
      var difference_ms = Math.abs(date1_ms - date2_ms);

      // Convert back to days and return
      return Math.round(difference_ms / ONE_DAY);
    }
  }
};

export default myMixin;

/*
 * Developed by Richard Malaschitz on 3/6/19 10:19 AM
 * Last modified 3/6/19 10:19 AM
 * Copyright (c) 2019. All right reserved.
 *
 */
