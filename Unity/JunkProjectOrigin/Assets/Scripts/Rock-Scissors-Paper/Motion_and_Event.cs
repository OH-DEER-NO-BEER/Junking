using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using static Hands_type;

public class Motion_and_Event : MonoBehaviour
{
    private bool flg = false;
    AudioSource audioSource;

    [HideInInspector]
    public bool mouse_flg = false;

    [HideInInspector]
    public bool selected = false;

    public Hands hand = Hands.Paper;

    private Quaternion rot;
    private Quaternion init_rot;

    [SerializeField]
    private GameObject canbas_obj;
    [SerializeField]
    private GameObject image_obj;

    private Canvas canvas;
    private RectTransform canvasrect;

    public Vector2 MousePos;

    private Image image;

    void OnMouseEnter()
    {
        if (mouse_flg)
        {
            audioSource.Play();

            flg = true;
        }
    }

    void OnMouseOver()
    {
        if (mouse_flg)
        {
            image.fillAmount += Time.deltaTime;
            if (image.fillAmount == 1)
            {
                audioSource.Stop();
                selected = true;
                Debug.Log("Selected!!!!");
                //Send
            }
        }
    }

    void OnMouseExit()
    {
        if (mouse_flg)
        {
            audioSource.Stop();

            flg = false;
            image.fillAmount = 0;
        }
    }

    // Start is called before the first frame update
    void Start()
    {
        audioSource = image_obj.GetComponent<AudioSource>();

        rot = Quaternion.AngleAxis(1, Vector3.right);
        init_rot = gameObject.transform.rotation;
        canvas = canbas_obj.GetComponent<Canvas>();
        canvasrect = canbas_obj.GetComponent<RectTransform>();
        image = image_obj.GetComponent<Image>();
        image.fillAmount = 0;
    }

    // Update is called once per frame
    void Update()
    {
        if (flg)
        {
            gameObject.transform.rotation *= rot;
            return;
        }

        RectTransformUtility.ScreenPointToLocalPointInRectangle(canvasrect,
                Input.mousePosition, canvas.worldCamera, out MousePos);

        image.GetComponent<RectTransform>().anchoredPosition = new Vector2(MousePos.x, MousePos.y);

        gameObject.transform.rotation = init_rot;
    }
}
