import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { ApiInterfaceService } from '../api-interface.service';

@Injectable({
  providedIn: 'root'
})
export class AuthguardService implements CanActivate {
  constructor(private apiService: ApiInterfaceService, private router: Router) { }
  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    if (!this.apiService.isAuthenticated()) {
      if (!this.router.url.includes('join-project')) {
        let invite_link = route.queryParams["invite-link"];
        console.log(invite_link)
        if (invite_link && invite_link.length != 0) {
          let redirect_url = `/login?invite-link=${invite_link}`
          this.router.navigateByUrl(redirect_url)
        }
        else
          this.router.navigate(['login']);
        return false;
      }
    }
    return true;
  }
}
