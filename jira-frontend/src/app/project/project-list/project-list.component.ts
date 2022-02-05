import { Component, OnInit } from '@angular/core';
import { ApiInterfaceService } from '../../api-interface.service';

export interface ProjectSettings {
  id: number,
  name: string;
  key: number;
  type: number;
  lead: string;
}

export var project_data: ProjectSettings[] = [
];

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.scss']
})
export class ProjectListComponent implements OnInit {
  displayedColumns: string[] = ['id', 'name', 'key', 'type', 'lead'];
  dataSource = project_data;

  constructor(private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
    this.apiService.getProjectList({'data':'data'})
    .subscribe((resp:any) => {
      this.dataSource = resp['response']
      project_data = this.dataSource
    })

  }
}
