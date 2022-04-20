import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';
@Component({
  selector: 'app-join',
  templateUrl: './join.component.html',
  styleUrls: ['./join.component.scss']
})
export class JoinComponent implements OnInit {
  isError: boolean = true
  isSuccess: boolean = true
  constructor(private route: ActivatedRoute, private apiService: ApiInterfaceService, private router: Router) {

  }

  ngOnInit(): void {
    let joinlink: string | null = this.route.snapshot.queryParams['invite-link']

    if (joinlink && joinlink.length != 0) {
      this.checkLink(joinlink)
    }
    else {
      this.router.navigate(['home','projects']);
    }
  }

  checkLink(joinlink: string | null): void {
    this.isError = false;
    this.isSuccess = false;
    this.apiService.joinProject(joinlink).subscribe((resp: any) => {
      if (!resp["status"]) {
        this.isError = true
      }
      else {
        this.isSuccess = true
      }
      setTimeout(() => { this.router.navigate(['home', 'projects']); }, 5000)
    })
  }
}
