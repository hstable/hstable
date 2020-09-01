export default function ({ store, redirect }) {
  // If the user is not authenticated
  if (!localStorage.token && !location.path.startsWith('/login')) {
    return redirect('/login')
  }
}
