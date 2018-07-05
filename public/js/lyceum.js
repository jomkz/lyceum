/*
 * Copyright 2018 Lyceum Developers
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

$( document ).ready(function() {
  $(".lyceum-date-relative").each(function(index) {
    var date = $(this).text();
    $(this).attr("title", moment(date).format('LLLL'));
    $(this).text(moment.utc(date).startOf('hour').fromNow());
  });

  $(".lyceum-filesize").each(function(index) {
    var bytes = $(this).text();
    $(this).attr("title", bytes + " bytes");
    $(this).text(Humanize.fileSize(bytes));
  });

  $("#linkDelete").click(function(e) {
    var url = e.currentTarget.href;
    if (confirm('Are you sure that you want to delete this item? This action is permanent.')) {
      var request = $.ajax({
        url:    url,
        method: "DELETE"
      });
      request.done(function( msg ) {
        window.location.replace("/");
      });
      request.fail(function( jqXHR, textStatus ) {
        console.log("delete API failed: " + jqXHR.status + " " + jqXHR.statusText);
      });
    }
    e.preventDefault();
  });
});
