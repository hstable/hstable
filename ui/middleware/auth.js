export default function ({ store, redirect }) {
  // If the user is not authenticated
  if (!localStorage.token && !location.pathname.startsWith('/login')) {
    return redirect('/login')
  }
}
