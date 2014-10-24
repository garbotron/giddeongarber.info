if (document.images)
{
  image_url = new Array();
  image_url[0]  = "treeview_about.png";
  image_url[1]  = "treeview_about_hot.png";
  image_url[2]  = "treeview_about_selected.png";
  image_url[3]  = "treeview_about_selected_hot.png";
  image_url[4]  = "treeview_resume.png";
  image_url[5]  = "treeview_resume_hot.png";
  image_url[6]  = "treeview_resume_selected.png";
  image_url[7]  = "treeview_resume_selected_hot.png";
  image_url[8]  = "treeview_projects.png";
  image_url[9]  = "treeview_projects_hot.png";
  image_url[10] = "treeview_projects_selected.png";
  image_url[11] = "treeview_projects_selected_hot.png";
  
  preload_image_object = new Image();
  var i;
  for (i = 0; i < 12; ++i) {
    preload_image_object.src = image_url[i];
  }
}

function MakeHot(elem) {
  elem.src = elem.getAttribute("id") + "_hot.png";
}

function MakeNormal(elem) {
  elem.src = elem.getAttribute("id") + ".png";
}