import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from '../api-interface.service';

@Component({
  selector: 'app-invite',
  templateUrl: './invite.component.html',
  styleUrls: ['./invite.component.scss']
})
export class InviteComponent implements OnInit {
  respMsg: string | null = null;
  errorMsg: string | null = null;
  respMap: any = {
    'ALREADY_PART_OF_PROJECT': 'This user is already part of the project',
    'INVITATION_SENT': 'This User has been Invited',
    'USER_ALREADY_INVITED': 'User is already invited',
    'ACTION_NOT_AUTHORIZED':'Only owners are allowed to add members to the project'
  }

  projectId: number = -1
  inviteForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.minLength(3)])
  })
  constructor(private route: ActivatedRoute, private apiService: ApiInterfaceService, private router: Router) {

  }

  ngOnInit(): void {
    let strProjectId = (this.route.snapshot.paramMap.get('projectId'))
    if (strProjectId) {
      this.projectId = parseInt(strProjectId)
    }
  }

  sendInvite() {
    let body = {
      email_id: this.inviteForm.get('email')?.value,
      project_id : this.projectId
    }

    this.apiService.sendInvite(body).subscribe((resp: any) => {
      if (resp["status"]) {
        let key = resp["message"]
        if (key in this.respMap) {
            this.respMsg = this.respMap[key]
        }

      } else {
        let key = resp["message"]
        if (key in this.respMap) {
          this.errorMsg = this.respMap[key]
        }
      }
    }, (err) => {

    })
    console.log(this.projectId, body.email_id)
  }

  isError() {
    return this.errorMsg && this.errorMsg.length != 0
  }

  isSuccess() {
    return this.respMsg && this.respMsg.length != 0
  }

}
