export default function ({ store, redirect }) {
  // If the user is not authenticated
  if (!localStorage.token) {
    return redirect('/login')
  }
}
